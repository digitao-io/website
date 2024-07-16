package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/gin-gonic/gin"
)

func PageList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		query, args, err := ctx.SqlBuilder.
			From("pages").
			Select(
				"key",
				"menu_name",
				"content_id",
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
			err = result.Scan(&page.Key, &page.MenuName, &page.ContentId)
			if err != nil {
				app.ResponseWithUnknownError(g, err)
				return
			}

			pages = append(pages, page)
		}

		app.ResponseWithData(g, pages)
	}
}
