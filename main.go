package main

import (
	"github.com/gin-gonic/gin"
	"todolist/api"
	"todolist/lib/gormLib"
	"todolist/utils/errorsUtils"
)

func main() {
	server := gin.Default()

	gormLib.MigrateDatabase()

	api.RegisterRouter(server)

	err := server.Run(":8080")
	errorsUtils.HandleErrorByPanic(err)
}
