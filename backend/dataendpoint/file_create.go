package dataendpoint

import (
	"context"
	"fmt"
	"time"

	"digitao.io/website/app"
	"digitao.io/website/datamodel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileCreate(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		data := datamodel.FileData{}
		err := g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
			return
		}

		newFileId := uuid.NewString()
		currentTime := time.Now()
		storedData := datamodel.File{
			FileIdentifier: datamodel.FileIdentifier{
				Id: &newFileId,
			},
			FileData: data,
			FileExtra: datamodel.FileExtra{
				CreatedAt: &currentTime,
			},
		}

		_, err = ctx.Database.Collection("files").
			InsertOne(context.Background(), storedData)
		if err != nil {
			app.ResponseWithUnknownError(g, err)
			return
		}

		app.ResponseWithData(g, gin.H{
			"newFileId":     newFileId,
			"fileUploadUrl": fmt.Sprintf("/data/file-upload/%s", newFileId),
		})
	}
}
