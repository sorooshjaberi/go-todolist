package gormLib

import (
	"booking/constants"
	"booking/lib/dotenvLib"
	"booking/models"
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
			dotenvLib.GetEnv("DB.HOST"), dotenvLib.GetEnv("DB.PORT"), dotenvLib.GetEnv("DB.USER"), dotenvLib.GetEnv("DB.PASSWORD"), dotenvLib.GetEnv("DB.DATABASE"))

		gormConfig := new(gorm.Config)

		returnedDb, err := gorm.Open(postgres.Open(dns), gormConfig)

		constants.HandleErrorByPanic(err)

		db = returnedDb

	})

	return db
}

func MigrateDatabase() {
	db := CreateConnection()
	err := db.AutoMigrate(new(models.User), new(models.Todo))

	constants.HandleErrorByPanic(err)
}
