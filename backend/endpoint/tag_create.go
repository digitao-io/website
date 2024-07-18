package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
)

func TagCreate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		data := model.Tag{}
		err := g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		query, args, err := ctx.SqlBuilder.
			Insert("tags").
			Rows(
				goqu.Record{
					"key":  data.Key,
					"name": data.Name,
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

		app.ResponseWithData(g, gin.H{"newTagKey": data.Key})
	}
}
