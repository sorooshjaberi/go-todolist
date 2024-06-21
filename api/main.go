package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitiateRouter(server *gin.Engine) {
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	})
}
