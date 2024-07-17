package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func FileEntryList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.FileEntrySearchParams{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		fileEntriesQuery := ctx.SqlBuilder.
			From("file_entries").
			Select(
				"id",
				"title",
				"mime_type",
				"size_in_bytes",
				"created_at",
			)

		whereClauses := []goqu.Expression{}
		if param.Query != nil && len(*param.Query) != 0 {
			whereClauses = append(
				whereClauses,
				goqu.Or(
					goqu.C("id").Eq(param.Query),
					goqu.L("MATCH (`title`) AGAINST (?)", param.Query),
				),
			)
		}
		if param.MimeType != nil && len(*param.MimeType) != 0 {
			whereClauses = append(
				whereClauses,
				goqu.C("mime_type").Eq(param.MimeType),
			)
		}
		fileEntriesQuery = fileEntriesQuery.
			Where(whereClauses...)

		if *param.Order == "ASC" {
			fileEntriesQuery = fileEntriesQuery.
				Order(goqu.C(*param.Sort).Asc())
		} else if *param.Order == "DESC" {
			fileEntriesQuery = fileEntriesQuery.
				Order(goqu.C(*param.Sort).Desc())
		}

		fileEntriesQuery = fileEntriesQuery.
			Offset(*param.Skip).
			Limit(*param.Take)

		query, args, err := fileEntriesQuery.ToSQL()
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

		app.ResponseWithData(g, fileEntries)
	}
}
