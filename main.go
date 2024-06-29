package main

import (
	"github.com/gin-gonic/gin"
	"todolist/api"
	"todolist/lib/gormLib"
	"todolist/middlewares"
	"todolist/utils/errorsUtils"
)

func main() {
	server := gin.Default()

	//todo use this inside a route group
	server.Use(middlewares.AuthGuard())

	gormLib.MigrateDatabase()

	api.InitiateRouter(server)

	err := server.Run(":8080")
	errorsUtils.HandleErrorByPanic(err)
}
