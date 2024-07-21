package dataendpoint

import (
	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(ctx *app.Context, r *gin.Engine) {
	r.POST("/data/tag-create", TagCreate(ctx))
	r.POST("/data/tag-list", TagList(ctx))
	r.POST("/data/tag-delete", TagDelete(ctx))

	r.POST("/data/article-create", ArticleCreate(ctx))
	r.POST("/data/article-get", ArticleGet(ctx))
	r.POST("/data/article-update", ArticleUpdate(ctx))
	r.POST("/data/article-list", ArticleList(ctx))
	r.POST("/data/article-delete", ArticleDelete(ctx))

	r.POST("/data/file-create", FileCreate(ctx))
	r.POST("/data/file-upload/:file-id", FileUpload(ctx))
	r.POST("/data/file-get", FileGet(ctx))
	r.GET("/data/file-download/:file-id", FileDownload(ctx))
	r.POST("/data/file-list", FileList(ctx))
	r.POST("/data/file-delete", FileDelete(ctx))
}
