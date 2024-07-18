package endpoint

import (
	"context"

	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func FileContentDelete(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		fileEntryId := g.Param("file-entry-id")

		err := ctx.Objstorage.RemoveObject(
			context.Background(),
			ctx.Configuration.Objstorage.Bucket,
			fileEntryId,
			minio.RemoveObjectOptions{},
		)

		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithOk(g)
	}
}
