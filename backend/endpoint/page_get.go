package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func PageGet(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.PageIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		query, args, err := ctx.SqlBuilder.
			From("pages").
			Select(
				"key",
				"menu_name",
				"article_id",
			).
			Where(
				goqu.C("key").Eq(param.Key),
			).
			ToSQL()
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		result, err := ctx.Database.Query(query, args...)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		pages := []model.Page{}
		for result.Next() {
			page := model.Page{}
			err = result.Scan(&page.Key, &page.MenuName, &page.ArticleId)
			if err != nil {
				app.ResponseWithUnknownError(g, err)
				return
			}

			pages = append(pages, page)
		}

		if len(pages) < 1 {
			app.ResponseWithEntityNotFound(g)
			return
		}

		app.ResponseWithData(g, pages[0])
	}
}
