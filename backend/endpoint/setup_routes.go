package endpoint

import (
	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(ctx *app.Context) *gin.Engine {
	r := gin.Default()

	r.GET("/do-health-check", DoHealthCheck(ctx))

	return r
}
