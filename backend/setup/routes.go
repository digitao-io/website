package setup

import (
	"digitao.io/website/app"
	"digitao.io/website/endpoint"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(ctx *app.Context) *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", endpoint.HealthCheck(ctx))

	r.POST("/tag-create", endpoint.TagCreate(ctx))
	r.POST("/tag-get", endpoint.TagGet(ctx))
	r.POST("/tag-list", endpoint.TagList(ctx))

	r.POST("/file-entry-create", endpoint.FileEntryCreate(ctx))

	r.POST("/page-create", endpoint.PageCreate(ctx))

	return r
}
