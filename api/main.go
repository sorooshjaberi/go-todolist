package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/api/auth"
	"todolist/api/todos"
	"todolist/constants"
	"todolist/middlewares"
)

func RegisterRouter(server *gin.Engine) {

	v1 := server.Group("/v1")

	{
		auth.RegisterRouter(v1)
	}
	{
		appRoute := v1.Group("/app")
		appRoute.Use(middlewares.AuthGuard())
		todos.RegisterRouter(appRoute)
	}

	v1.GET("/", func(ctx *gin.Context) {
		fmt.Println(ctx.Get(constants.Keys.RequestUser))
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	})

}
