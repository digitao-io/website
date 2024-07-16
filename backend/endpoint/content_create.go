package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ContentCreate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		data := model.ContentData{}
		err := g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		contentId := uuid.NewString()

		query, args, err := ctx.SqlBuilder.
			Insert("contents").
			Rows(
				goqu.Record{
					"id":        contentId,
					"type":      data.Type,
					"title":     data.Title,
					"createdAt": data.CreatedAt,
					"updatedAt": data.UpdatedAt,
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

		if len(*data.TagKeys) != 0 {
			contentTagLinks := []goqu.Record{}

			for _, tagKey := range *data.TagKeys {
				contentTagLinks = append(
					contentTagLinks,
					goqu.Record{
						"content_id": contentId,
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

		app.ResponseWithData(g, gin.H{"newContentId": contentId})
	}
}
