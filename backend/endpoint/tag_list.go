package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/gin-gonic/gin"
)

func TagList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		query, args, err := ctx.SqlBuilder.
			From("tags").
			Select(
				"key",
				"name",
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

		app.ResponseWithData(g, tags)
	}
}
