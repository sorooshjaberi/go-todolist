package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/api/auth"
	"todolist/constants"
)

func InitiateRouter(server *gin.Engine) {

	v1 := server.Group("/v1")

	auth.InitiateRouter(v1)

	v1.GET("/", func(ctx *gin.Context) {
		fmt.Println(ctx.Get(constants.Keys.RequestUser))
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	})

}
