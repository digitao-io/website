package siteendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/sitemodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ResourceUpdate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		param := sitemodel.ResourceIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid URL query")
			return
		}

		data := sitemodel.Resource{}
		err = g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid request body")
			return
		}

		result, err := ctx.Database.Collection("resources").
			UpdateOne(
				context.Background(),
				bson.D{{Key: "key", Value: *param.Key}},
				bson.D{{Key: "$set", Value: data}},
			)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		if result.ModifiedCount == 0 {
			app.ResponseWithEntityNotFound(g)
			return
		}

		app.ResponseWithData(g, gin.H{"updatedResourceKey": data.Key})
	}
}
