package siteendpoint

import (
	"digitao.io/website/app"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(ctx *app.Context, r *gin.Engine) {
	r.GET("/site/health-check", HealthCheck(ctx))

	r.POST("/site/user-login", UserLogin(ctx))
	r.POST("/site/user-auth-challenge", UserAuthChallenge(ctx))
	r.POST("/site/user-logout", UserLogout(ctx))

	r.POST("/site/page-create", PageCreate(ctx))
	r.POST("/site/page-get", PageGet(ctx))
	r.POST("/site/page-list", PageList(ctx))
	r.POST("/site/page-update", PageUpdate(ctx))
	r.POST("/site/page-delete", PageDelete(ctx))

	r.POST("/site/resource-create", ResourceCreate(ctx))
	r.POST("/site/resource-get", ResourceGet(ctx))
	r.POST("/site/resource-list", ResourceList(ctx))
	r.POST("/site/resource-update", ResourceUpdate(ctx))
	r.POST("/site/resource-delete", ResourceDelete(ctx))
}
