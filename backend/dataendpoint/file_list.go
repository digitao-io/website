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

func FileList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		param := datamodel.FileSearchParams{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid URL query")
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

		if param.MimeType != nil && len(*param.MimeType) != 0 {
			filter = append(
				filter,
				bson.E{
					Key:   "mimeType",
					Value: *param.MimeType,
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

		results := []datamodel.File{}
		cursor, err := ctx.Database.Collection("files").
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
