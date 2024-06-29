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

	//public routes
	{
		auth.RegisterRouter(v1)
	}
	//protectedRoutes
	{
		protectedRoutes := v1.Group("/")
		protectedRoutes.Use(middlewares.AuthGuard())
		todos.RegisterRouter(protectedRoutes)
	}

	v1.GET("/", func(ctx *gin.Context) {
		fmt.Println(ctx.Get(constants.Keys.RequestUser))
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	})

}
