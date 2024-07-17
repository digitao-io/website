package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
)

func PageUpdate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.PageIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		data := model.Page{}
		err = g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		query, args, err := ctx.SqlBuilder.
			Update("pages").
			Set(
				goqu.Record{
					"key":        data.Key,
					"menu_name":  data.MenuName,
					"content_id": data.ContentId,
				},
			).
			Where(
				goqu.C("key").Eq(param.Key),
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

		app.ResponseWithOk(g)
	}
}
