package dataendpoint

import (
	"context"
	"time"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ArticleCreate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		data := datamodel.ArticleData{}
		err := g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		newArticleId := uuid.NewString()
		currentTime := time.Now()
		storedData := datamodel.Article{
			ArticleIdentifier: datamodel.ArticleIdentifier{
				Id: &newArticleId,
			},
			ArticleData: data,
			ArticleExtra: datamodel.ArticleExtra{
				CreatedAt: &currentTime,
				UpdatedAt: &currentTime,
			},
		}

		_, err = ctx.Database.Collection("articles").
			InsertOne(context.Background(), storedData)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithData(g, gin.H{"newArticleId": newArticleId})
	}
}
