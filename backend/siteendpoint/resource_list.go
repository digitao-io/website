package siteendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/sitemodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ResourceList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		param := sitemodel.ResourceSearchParams{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid URL query")
			return
		}

		filter := bson.D{}
		if param.Type != nil && len(*param.Type) != 0 {
			filter = append(filter, bson.E{Key: "type", Value: *param.Type})
		}

		results := []sitemodel.Resource{}
		cursor, err := ctx.Database.Collection("resources").
			Find(context.Background(), filter)
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
