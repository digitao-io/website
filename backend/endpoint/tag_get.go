package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func TagGet(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.TagIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		query, args, err := ctx.SqlBuilder.
			From("tags").
			Select(
				"key",
				"name",
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

		tags := []model.Tag{}
		for result.Next() {
			tag := model.Tag{}
			err = result.Scan(&tag.Key, &tag.Name)
			if err != nil {
				app.ResponseWithUnknownError(g, err)
				return
			}

			tags = append(tags, tag)
		}

		if len(tags) < 1 {
			app.ResponseWithEntityNotFound(g)
			return
		}

		app.ResponseWithData(g, tags[0])
	}
}
