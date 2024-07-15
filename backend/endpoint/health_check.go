package endpoint

import (
	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
)

func DoHealthCheck(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		ResponseWithOk(g)
	}
}
