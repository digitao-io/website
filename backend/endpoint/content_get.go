package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func ContentGet(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.ContentIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		query, args, err := ctx.SqlBuilder.
			From("contents").
			LeftJoin(
				goqu.T("content_tag_links"),
				goqu.On(goqu.I("contents.id").Eq(goqu.I("content_tag_links.content_id"))),
			).
			LeftJoin(
				goqu.T("tags"),
				goqu.On(goqu.I("tags.key").Eq(goqu.I("content_tag_links.tag_key"))),
			).
			Select(
				"contents.id",
				"contents.type",
				"contents.title",
				"contents.created_at",
				"contents.updated_at",
				"contents.summary",
				"contents.content",
				"tags.key",
				"tags.name",
			).
			Where(
				goqu.I("contents.id").Eq(param.Id),
			).
			ToSQL()
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		contentMap := map[string]model.Content{}
		contentTagMap := map[string]map[string]model.Tag{}

		results, err := ctx.Database.Query(query, args...)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		for results.Next() {
			content := model.Content{}
			tag := model.Tag{}

			results.Scan(
				&content.Id,
				&content.Type,
				&content.Title,
				&content.CreatedAt,
				&content.UpdatedAt,
				&content.Summary,
				&content.Content,
				&tag.Key,
				&tag.Name,
			)

			if content.Id != nil {
				if _, ok := contentMap[*content.Id]; !ok {
					contentMap[*content.Id] = content
				}
			}

			if tag.Key != nil {
				if _, ok := contentTagMap[*content.Id]; !ok {
					contentTagMap[*content.Id] = map[string]model.Tag{}
				}
				if _, ok := contentTagMap[*content.Id][*tag.Key]; !ok {
					contentTagMap[*content.Id][*tag.Key] = tag
				}
			}
		}

		contents := []model.Content{}

		for _, content := range contentMap {
			tagKeys := []string{}
			content.TagKeys = &tagKeys

			for _, tag := range contentTagMap[*content.Id] {
				*content.TagKeys = append(*content.TagKeys, *tag.Key)
			}
			contents = append(contents, content)
		}

		if len(contents) < 1 {
			app.ResponseWithEntityNotFound(g)
			return
		}

		app.ResponseWithData(g, contents[0])
	}
}
