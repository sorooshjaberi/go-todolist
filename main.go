package main

import (
	"booking/api"
	"booking/lib/gormLib"
	"booking/middlewares"
	"booking/utils/errorsUtils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(middlewares.AuthGuard())

	gormLib.MigrateDatabase()

	api.InitiateRouter(server)

	err := server.Run(":8080")
	errorsUtils.HandleErrorByPanic(err)
}
