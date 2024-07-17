package endpoint

import (
	"time"

	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
)

func ContentUpdate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.ContentIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		data := model.ContentData{}
		err = g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")

		query, args, err := ctx.SqlBuilder.
			Update("contents").
			Set(
				goqu.Record{
					"type":      data.Type,
					"title":     data.Title,
					"updated_at": currentTime,
					"summary":   data.Summary,
					"content":   data.Content,
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

		query, args, err = ctx.SqlBuilder.
			Delete("content_tag_links").
			Where(
				goqu.C("content_id").Eq(param.Id),
			).
			ToSQL()
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		_, err = ctx.Database.Exec(query, args...)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		if len(*data.TagKeys) != 0 {
			contentTagLinks := []goqu.Record{}

			for _, tagKey := range *data.TagKeys {
				contentTagLinks = append(
					contentTagLinks,
					goqu.Record{
						"content_id": param.Id,
						"tag_key":    tagKey,
					},
				)
			}

			query, args, err := ctx.SqlBuilder.
				Insert("content_tag_links").
				Rows(contentTagLinks).
				ToSQL()
			if err != nil {
				app.ResponseWithUnknownError(g, err)
				return
			}

			_, err = ctx.Database.Exec(query, args...)
			if err != nil {
				app.ResponseWithUnknownError(g, err)
				return
			}
		}

		app.ResponseWithOk(g)
	}
}
