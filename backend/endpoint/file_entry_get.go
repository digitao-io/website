package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func FileEntryGet(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.FileEntryIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		query, args, err := ctx.SqlBuilder.
			From("file_entries").
			Select(
				"id",
				"title",
				"mime_type",
				"size_in_bytes",
			).
			Where(
				goqu.C("id").Eq(param.Id),
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

		app.ResponseWithData(g, fileEntries[0])
	}
}
