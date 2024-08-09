package siteendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/sitemodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PageGet(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := sitemodel.PageIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid URL query")
			return
		}

		result := sitemodel.Page{}
		err = ctx.Database.Collection("pages").
			FindOne(context.Background(), bson.D{{Key: "key", Value: *param.Key}}).
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
