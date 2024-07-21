package siteendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/sitemodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func PageList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		results := []sitemodel.Page{}
		cursor, err := ctx.Database.Collection("pages").
			Find(context.Background(), bson.D{})
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
