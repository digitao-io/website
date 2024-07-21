package siteendpoint

import (
	"crypto/sha256"
	"encoding/base64"
	"time"

	"digitao.io/website/app"
	"digitao.io/website/sitemodel"
	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *app.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		data := sitemodel.LoginRequest{}
		err := g.ShouldBindJSON(&data)
		if err != nil {
			app.ResponseWithParseError(g, "Cannot parse request body")
		}

		for _, user := range ctx.Configuration.Users {
			if user.Username != *data.Username {
				continue
			}

			hashBytes := sha256.Sum256([]byte(*data.Password))
			hashStr := base64.URLEncoding.EncodeToString(hashBytes[:])

			if user.PasswordHash != hashStr {
				continue
			}

			jwt := app.SignJwt(ctx.Configuration)
			maxAge, err := time.ParseDuration(ctx.Configuration.Jwt.ExpireIn)
			if err != nil {
				app.ResponseWithUnknownError(g, err)
				return
			}

			g.SetCookie("jwt", jwt, int(maxAge.Seconds()), "/", ctx.Configuration.Domain, false, true)
			app.ResponseWithOk(g)
			return
		}

		app.ResponseWithAuthenticationFailed(g)
	}
}
