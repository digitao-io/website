package setup

import (
	"digitao.io/website/app"
	"digitao.io/website/endpoint"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(ctx *app.Context) *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", endpoint.HealthCheck(ctx))

	r.POST("/content-create", endpoint.ContentCreate(ctx))
	r.POST("/content-delete", endpoint.ContentDelete(ctx))
	r.POST("/content-get", endpoint.ContentGet(ctx))

	r.POST("/tag-create", endpoint.TagCreate(ctx))
	r.POST("/tag-get", endpoint.TagGet(ctx))
	r.POST("/tag-list", endpoint.TagList(ctx))
	r.POST("/tag-delete", endpoint.TagDelete(ctx))

	r.POST("/file-entry-create", endpoint.FileEntryCreate(ctx))
	r.POST("/file-entry-get", endpoint.FileEntryGet(ctx))
	r.POST("/file-entry-delete", endpoint.FileEntryDelete(ctx))

	r.POST("/page-create", endpoint.PageCreate(ctx))
	r.POST("/page-get", endpoint.PageGet(ctx))
	r.POST("/page-list", endpoint.PageList(ctx))
	r.POST("/page-update", endpoint.PageUpdate(ctx))
	r.POST("/page-delete", endpoint.PageDelete(ctx))

	return r
}
