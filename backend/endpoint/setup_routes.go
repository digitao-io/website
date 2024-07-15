package endpoint

import (
	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(ctx *app.Context) *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", HealthCheck(ctx))

	return r
}
