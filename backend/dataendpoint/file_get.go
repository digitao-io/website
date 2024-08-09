package dataendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FileGet(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		param := datamodel.FileIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid URL query")
			return
		}

		result := datamodel.File{}
		err = ctx.Database.Collection("files").
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
