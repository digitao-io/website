package endpoint

import (
	"context"
	"net/http"

	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func FileContentDownload(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		fileEntryId := g.Param("file-entry-id")

		query, args, err := ctx.SqlBuilder.
			From("file_entries").
			Select(
				"id",
				"title",
				"mime_type",
				"size_in_bytes",
				"created_at",
			).
			Where(
				goqu.C("id").Eq(fileEntryId),
			).
			ToSQL()
		if err != nil {
			g.AbortWithStatus(http.StatusNotFound)
			return
		}

		results, err := ctx.Database.Query(query, args...)
		if err != nil {
			g.AbortWithStatus(http.StatusNotFound)
			return
		}

		fileEntries := []model.FileEntry{}
		for results.Next() {
			fileEntry := model.FileEntry{}
			err = results.Scan(
				&fileEntry.Id,
				&fileEntry.Title,
				&fileEntry.MimeType,
				&fileEntry.SizeInBytes,
				&fileEntry.CreatedAt,
			)
			if err != nil {
				g.AbortWithStatus(http.StatusNotFound)
				return
			}

			fileEntries = append(fileEntries, fileEntry)
		}

		if len(fileEntries) < 1 {
			g.AbortWithStatus(http.StatusNotFound)
			return
		}

		fileEntry := fileEntries[0]

		fileReader, err := ctx.Objstorage.GetObject(
			context.Background(),
			ctx.Configuration.Objstorage.Bucket,
			*fileEntry.Id,
			minio.GetObjectOptions{},
		)
		if err != nil {
			g.AbortWithStatus(http.StatusNotFound)
			return
		}

		g.DataFromReader(
			http.StatusOK,
			*fileEntry.SizeInBytes,
			*fileEntry.MimeType,
			fileReader,
			map[string]string{},
		)
	}
}
