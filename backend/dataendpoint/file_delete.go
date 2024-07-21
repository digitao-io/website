package dataendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson"
)

func FileDelete(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		param := datamodel.FileIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		_, err = ctx.Database.Collection("files").
			DeleteOne(context.Background(), bson.D{{Key: "_id", Value: *param.Id}})
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		err = ctx.Objstorage.RemoveObject(
			context.Background(),
			ctx.Configuration.Objstorage.Bucket,
			*param.Id,
			minio.RemoveObjectOptions{},
		)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithOk(g)
	}
}
