package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
)

func PageCreate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		data := model.Page{}
		err := g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		query, args, err := ctx.SqlBuilder.
			Insert("pages").
			Rows(
				goqu.Record{
					"key":        data.Key,
					"menu_name":  data.MenuName,
					"article_id": data.ArticleId,
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

		app.ResponseWithData(g, gin.H{"newPageKey": data.Key})
	}
}
