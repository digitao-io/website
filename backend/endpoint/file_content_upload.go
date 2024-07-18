package endpoint

import (
	"context"

	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func FileContentUpload(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		fileEntryId := g.Param("file-entry-id")

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
			app.ResponseWithUnknownError(g, err)
			return
		}

		results, err := ctx.Database.Query(query, args...)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
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
				app.ResponseWithUnknownError(g, err)
				return
			}

			fileEntries = append(fileEntries, fileEntry)
		}

		if len(fileEntries) < 1 {
			app.ResponseWithEntityNotFound(g)
			return
		}

		fileEntry := fileEntries[0]

		_, err = ctx.Objstorage.PutObject(
			context.Background(),
			ctx.Configuration.Objstorage.Bucket,
			*fileEntry.Id,
			fileReader,
			*fileEntry.SizeInBytes,
			minio.PutObjectOptions{
				ContentType: *fileEntry.MimeType,
			},
		)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithOk(g)
	}
}
