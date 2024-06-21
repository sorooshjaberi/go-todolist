package gorm

import (
	"booking/lib/dotenv"
	"booking/models"
	"booking/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func CreateConnection() *gorm.DB {
	once.Do(func() {
		dns := fmt.Sprintf("host=%v port=%v user=%v password=%v database=%v",
			dotenv.GetEnv("DB.HOST"), dotenv.GetEnv("DB.PORT"), dotenv.GetEnv("DB.USER"), dotenv.GetEnv("DB.PASSWORD"), dotenv.GetEnv("DB.DATABASE"))

		gormConfig := new(gorm.Config)

		returnedDb, err := gorm.Open(postgres.Open(dns), gormConfig)

		utils.HandleErrorByPanic(err)

		db = returnedDb

	})

	return db
}

func MigrateDatabase() {
	db := CreateConnection()
	err := db.AutoMigrate(new(models.User), new(models.Todo))

	utils.HandleErrorByPanic(err)
}
