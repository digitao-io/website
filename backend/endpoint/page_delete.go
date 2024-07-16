package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func PageDelete(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.PageIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		query, args, err := ctx.SqlBuilder.
			Delete("pages").
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
