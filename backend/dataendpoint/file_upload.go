package dataendpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FileUpload(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		fileId := g.Param("file-id")

		file, err := g.FormFile("file")
		if err != nil {
			app.ResponseWithParseError(g, "Cannot read file")
			return
		}
		fileReader, err := file.Open()
		if err != nil {
			app.ResponseWithParseError(g, "Cannot read file")
			return
		}
		defer fileReader.Close()

		result := datamodel.File{}
		err = ctx.Database.Collection("files").
			FindOne(context.Background(), bson.D{{Key: "_id", Value: fileId}}).
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

		_, err = ctx.Objstorage.PutObject(
			context.Background(),
			ctx.Configuration.Objstorage.Bucket,
			*result.Id,
			fileReader,
			*result.SizeInBytes,
			minio.PutObjectOptions{
				ContentType: *result.MimeType,
			},
		)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithOk(g)
	}
}
