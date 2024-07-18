package app

import "github.com/gin-gonic/gin"

func CheckPermission(g *gin.Context, config *Configuration) bool {
	jwt, err := g.Cookie("jwt")
	if err != nil {
		return false
	}

	return CheckJwt(config, jwt)
}
