package api

import (
	"booking/api/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitiateRouter(server *gin.Engine) {

	v1 := server.Group("/v1")

	auth.InitiateRouter(v1)

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	})

}
