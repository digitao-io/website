package siteendpoint

import (
	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
)

func UserAuthChallenge(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		if !app.CheckPermission(g, ctx.Configuration) {
			app.ResponseWithAuthenticationFailed(g)
			return
		}

		app.ResponseWithOk(g)
	}
}
