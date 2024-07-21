package dataendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TagList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		results := []datamodel.Tag{}
		cursor, err := ctx.Database.Collection("tags").
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
