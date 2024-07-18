package setup

import (
	"digitao.io/website/app"
	"digitao.io/website/endpoint"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(ctx *app.Context) *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", endpoint.HealthCheck(ctx))

	r.POST("/user-login", endpoint.UserLogin(ctx))
	r.POST("/user-logout", endpoint.UserLogout(ctx))

	r.POST("/content-create", endpoint.ContentCreate(ctx))
	r.POST("/content-delete", endpoint.ContentDelete(ctx))
	r.POST("/content-get", endpoint.ContentGet(ctx))
	r.POST("/content-list", endpoint.ContentList(ctx))
	r.POST("/content-update", endpoint.ContentUpdate(ctx))

	r.POST("/tag-create", endpoint.TagCreate(ctx))
	r.POST("/tag-get", endpoint.TagGet(ctx))
	r.POST("/tag-list", endpoint.TagList(ctx))
	r.POST("/tag-delete", endpoint.TagDelete(ctx))

	r.POST("/file-content-upload/:file-entry-id", endpoint.FileContentUpload(ctx))
	r.POST("/file-content-delete/:file-entry-id", endpoint.FileContentDelete(ctx))
	r.GET("/file-content-download/:file-entry-id", endpoint.FileContentDownload(ctx))

	r.POST("/file-entry-create", endpoint.FileEntryCreate(ctx))
	r.POST("/file-entry-get", endpoint.FileEntryGet(ctx))
	r.POST("/file-entry-list", endpoint.FileEntryList(ctx))
	r.POST("/file-entry-delete", endpoint.FileEntryDelete(ctx))

	r.POST("/page-create", endpoint.PageCreate(ctx))
	r.POST("/page-get", endpoint.PageGet(ctx))
	r.POST("/page-list", endpoint.PageList(ctx))
	r.POST("/page-update", endpoint.PageUpdate(ctx))
	r.POST("/page-delete", endpoint.PageDelete(ctx))

	return r
}
