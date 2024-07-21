package dataendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ArticleGet(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := datamodel.ArticleIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		result := datamodel.Article{}
		err = ctx.Database.Collection("articles").
			FindOne(context.Background(), bson.D{{Key: "_id", Value: *param.Id}}).
			Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				app.ResponseWithEntityNotFound(g)
				return
			} else {
				app.ResponseWithUnknownError(g, err)
				return
			}
		}

		app.ResponseWithData(g, result)
	}
}
