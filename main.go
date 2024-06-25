package main

import (
	"booking/api"
	"booking/constants"
	"booking/lib/gormLib"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	gormLib.MigrateDatabase()

	api.InitiateRouter(server)

	err := server.Run(":8080")
	constants.HandleErrorByPanic(err)
}
