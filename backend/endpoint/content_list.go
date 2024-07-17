package endpoint

import (
	"digitao.io/website/app"
	"digitao.io/website/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func ContentList(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		param := model.ContentSearchParams{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request query parameters")
			return
		}

		contentIdQuery := ctx.SqlBuilder.
			From("contents").
			LeftJoin(
				goqu.T("content_tag_links"),
				goqu.On(goqu.Ex{"contents.id": goqu.I("content_tag_links.content_id")}),
			).
			LeftJoin(
				goqu.T("tags"),
				goqu.On(goqu.Ex{"tags.key": goqu.I("content_tag_links.tag_key")}),
			)

		selectColumns := []interface{}{
			goqu.I("contents.id").As("id"),
			goqu.I("contents.created_at").As("created_at"),
			goqu.I("contents.updated_at").As("updated_at"),
		}
		if param.Query != nil && len(*param.Query) != 0 {
			selectColumns = append(
				selectColumns,
				goqu.MAX(
					goqu.L("MATCH (`contents`.`title`, `contents`.`summary`) AGAINST (?)", param.Query),
				).As("score"),
			)
		} else {
			selectColumns = append(
				selectColumns,
				goqu.V(0).As("score"),
			)
		}
		contentIdQuery = contentIdQuery.
			Select(selectColumns...)

		whereClauses := []goqu.Expression{}
		if param.Query != nil && len(*param.Query) != 0 {
			whereClauses = append(
				whereClauses,
				goqu.Or(
					goqu.L("contents.id").Eq(param.Query),
					goqu.L("MATCH (`contents`.`title`, `contents`.`summary`) AGAINST (?)", param.Query),
				),
			)
		}
		if param.Type != nil && len(*param.Type) != 0 {
			whereClauses = append(
				whereClauses,
				goqu.L("contents.type").Eq(param.Type),
			)
		}
		if param.Tags != nil && len(*param.Tags) != 0 {
			whereClauses = append(
				whereClauses,
				goqu.L("tags.key").In(param.Tags),
			)
		}
		contentIdQuery = contentIdQuery.
			Where(whereClauses...)

		contentIdQuery = contentIdQuery.
			GroupBy(goqu.C("id"))

		if param.Tags != nil && len(*param.Tags) != 0 {
			contentIdQuery = contentIdQuery.
				Having(goqu.COUNT(1).Eq(len(*param.Tags)))
		}

		if *param.Order == "ASC" {
			contentIdQuery = contentIdQuery.
				Order(goqu.C(*param.Sort).Asc())
		} else if *param.Order == "DESC" {
			contentIdQuery = contentIdQuery.
				Order(goqu.C(*param.Sort).Desc())
		}

		contentIdQuery = contentIdQuery.
			Offset(*param.Skip).
			Limit(*param.Take)

		query, args, err := ctx.SqlBuilder.
			From(
				contentIdQuery.As("r"),
			).
			LeftJoin(
				goqu.T("contents"),
				goqu.On(goqu.I("r.id").Eq(goqu.I("contents.id"))),
			).
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
			ToSQL()
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		contentIds := []string{}
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
					contentIds = append(contentIds, *content.Id)
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

		for _, contentId := range contentIds {
			content := contentMap[contentId]

			tagKeys := []string{}
			content.TagKeys = &tagKeys

			for _, tag := range contentTagMap[*content.Id] {
				*content.TagKeys = append(*content.TagKeys, *tag.Key)
			}
			contents = append(contents, content)
		}

		app.ResponseWithData(g, contents)
	}
}
