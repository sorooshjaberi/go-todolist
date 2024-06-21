package main

import (
	"booking/api"
	"booking/lib/gorm"
	"booking/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	gorm.MigrateDatabase()

	api.InitiateRouter(server)

	err := server.Run(":8080")
	utils.HandleErrorByPanic(err)
}
