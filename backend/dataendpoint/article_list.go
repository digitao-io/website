package dataendpoint

import (
	"context"
	"fmt"
	"strings"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ArticleList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := datamodel.ArticleSearchParams{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		filter := bson.D{}

		if param.Query != nil && len(*param.Query) != 0 {
			queryString := ""
			for i, word := range strings.Split(*param.Query, " ") {
				if len(word) == 0 {
					continue
				}
				if i != 0 {
					queryString += " "
				}
				queryString += fmt.Sprintf("\"%s\"", word)
			}

			filter = append(
				filter,
				bson.E{
					Key:   "$text",
					Value: bson.D{{Key: "$search", Value: queryString}},
				},
			)
		}

		if param.Tags != nil && len(*param.Tags) != 0 {
			filter = append(
				filter,
				bson.E{
					Key:   "tagKeys",
					Value: bson.D{{Key: "$all", Value: *param.Tags}},
				},
			)
		}

		var sort bson.D

		if *param.Sort == "score" {
			sort = bson.D{{
				Key:   "score",
				Value: bson.D{{Key: "$meta", Value: "textScore"}},
			}}
		} else {
			var order int
			if *param.Order == "ASC" {
				order = 1
			} else {
				order = -1
			}
			sort = bson.D{{Key: *param.Sort, Value: order}}
		}

		results := []datamodel.Article{}
		cursor, err := ctx.Database.Collection("articles").
			Find(
				context.Background(),
				filter,
				options.Find().
					SetSort(sort).
					SetSkip(*param.Skip).
					SetLimit(*param.Take),
			)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		err = cursor.All(context.Background(), &results)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithData(g, results)
	}
}
