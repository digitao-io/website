package dataendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TagDelete(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		param := datamodel.TagIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid URL query")
			return
		}

		_, err = ctx.Database.Collection("tags").
			DeleteOne(context.Background(), bson.D{{Key: "key", Value: *param.Key}})
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithOk(g)
	}
}
