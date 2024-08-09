package dataendpoint

import (
	"context"
	"time"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ArticleUpdate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		param := datamodel.ArticleIdentifier{}
		err := g.ShouldBindQuery(&param)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid URL query")
			return
		}

		data := datamodel.ArticleData{}
		err = g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithValidationFailed(g, "Invalid request body")
			return
		}

		currentTime := time.Now()
		storedData := datamodel.Article{
			ArticleData: data,
			ArticleExtra: datamodel.ArticleExtra{
				UpdatedAt: &currentTime,
			},
		}

		result, err := ctx.Database.Collection("articles").
			UpdateOne(
				context.Background(),
				bson.D{{Key: "_id", Value: *param.Id}},
				bson.D{{Key: "$set", Value: storedData}},
			)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		if result.ModifiedCount == 0 {
			app.ResponseWithEntityNotFound(g)
			return
		}

		app.ResponseWithData(g, gin.H{"updatedArticleId": param.Id})
	}
}
