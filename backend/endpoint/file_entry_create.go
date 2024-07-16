package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileEntryCreate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		data := model.FileEntryData{}
		err := g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithValidationFailed(g, err.Error())
			return
		}

		fileEntryId := uuid.NewString()

		query, args, err := ctx.SqlBuilder.
			Insert("file_entries").
			Rows(
				goqu.Record{
					"id":            fileEntryId,
					"title":         data.Title,
					"mime_type":     data.MimeType,
					"size_in_bytes": data.SizeInBytes,
				},
			).ToSQL()
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		_, err = ctx.Database.Exec(query, args...)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithData(g, gin.H{"newFileEntryId": fileEntryId})
	}
}
