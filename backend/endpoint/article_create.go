package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ArticleCreate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		data := model.ArticleData{}
		err := g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		articleId := uuid.NewString()

		query, args, err := ctx.SqlBuilder.
			Insert("articles").
			Rows(
				goqu.Record{
					"id":        articleId,
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

		if len(data.TagKeys) != 0 {
			tageRelationships := []goqu.Record{}

			for _, tageKey := range data.TagKeys {
				tageRelationships = append(tageRelationships, goqu.Record{
					"article_id": articleId,
					"tag_key":    tageKey,
				})
			}
			query, args, err := ctx.SqlBuilder.
				Insert("article_tag_links").
				Rows(
					tageRelationships,
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

			app.ResponseWithData(g, gin.H{"newArticleId": articleId})
		}

	}
}
