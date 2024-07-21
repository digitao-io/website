package siteendpoint

import (
	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		app.ResponseWithOk(g)
	}
}
