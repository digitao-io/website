package endpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const STATUS_OK = "OK"
const STATUS_VALIDATION_FAILED = "VALIDATION_FAILED"
const STATUS_AUTHENTICATION_FAILED = "AUTHENTICATION_FAILED"
const STATUS_ENTITY_NOT_FOUND = "ENTITY_NOT_FOUND"
const STATUS_UNKNOWN_ERROR = "UNKNOWN_ERROR"

func ResponseWithOk(g *gin.Context) {
	g.JSON(200, gin.H{
		"status": STATUS_OK,
	})
}

func ResponseWithData(g *gin.Context, data any) {
	g.JSON(200, gin.H{
		"status": STATUS_OK,
		"data":   data,
	})
}

func ResponseWithValidationFailed(g *gin.Context, err string) {
	g.JSON(200, gin.H{
		"status": STATUS_VALIDATION_FAILED,
		"error":  err,
	})
}

func ResponseWithAuthenticationFailed(g *gin.Context) {
	g.JSON(200, gin.H{
		"status": STATUS_AUTHENTICATION_FAILED,
	})
}

func ResponseWithEntityNotFound(g *gin.Context) {
	g.JSON(200, gin.H{
		"status": STATUS_ENTITY_NOT_FOUND,
	})
}

func ResponseWithUnknownError(g *gin.Context, err error) {
	g.JSON(200, gin.H{
		"status": STATUS_UNKNOWN_ERROR,
		"error":  err.Error(),
	})

	fmt.Println(err)
}
