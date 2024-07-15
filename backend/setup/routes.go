package setup

import (
	"digitao.io/website/app"
	"digitao.io/website/endpoint"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(ctx *app.Context) *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", endpoint.HealthCheck(ctx))

	return r
}
