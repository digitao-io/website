package siteendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/sitemodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func PageUpdate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		param := sitemodel.PageIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		data := sitemodel.Page{}
		err = g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		result, err := ctx.Database.Collection("pages").
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

		app.ResponseWithData(g, gin.H{"updatedPageKey": data.Key})
	}
}
