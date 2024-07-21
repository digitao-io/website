package dataendpoint

import (
	"context"
	"net/http"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FileDownload(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		fileId := g.Param("file-id")

		result := datamodel.File{}
		err := ctx.Database.Collection("files").
			FindOne(context.Background(), bson.D{{Key: "_id", Value: fileId}}).
			Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				g.AbortWithStatus(http.StatusNotFound)
				return
			} else {
				g.AbortWithStatus(http.StatusNotFound)
				return
			}
		}

		fileReader, err := ctx.Objstorage.GetObject(
			context.Background(),
			ctx.Configuration.Objstorage.Bucket,
			*result.Id,
			minio.GetObjectOptions{},
		)
		if err != nil {
			g.AbortWithStatus(http.StatusNotFound)
			return
		}

		g.DataFromReader(
			http.StatusOK,
			*result.SizeInBytes,
			*result.MimeType,
			fileReader,
			map[string]string{},
		)
	}
}
