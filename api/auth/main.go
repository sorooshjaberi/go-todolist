package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitiateRouter(server *gin.RouterGroup) {
	authRouter := server.Group("/auth")

	authRouter.POST("/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "auth login route",
		})
	})
}
