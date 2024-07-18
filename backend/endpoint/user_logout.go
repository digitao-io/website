package endpoint

import (
	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
)

func UserLogout(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		g.SetCookie("jwt", "", 0, "/", ctx.Configuration.Domain, false, true)
		app.ResponseWithOk(g)
	}
}
