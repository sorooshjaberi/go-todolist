package gormLib

import (
	"booking/constants"
	"booking/lib/dotenvLib"
	"booking/models"
	"booking/utils/errorsUtils"
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
			dotenvLib.GetEnv(constants.EnvKeys.DBHost),
			dotenvLib.GetEnv(constants.EnvKeys.DBPort),
			dotenvLib.GetEnv(constants.EnvKeys.DBUser),
			dotenvLib.GetEnv(constants.EnvKeys.DBPass),
			dotenvLib.GetEnv(constants.EnvKeys.DBDatabase))

		fmt.Println(dns)
		gormConfig := new(gorm.Config)

		returnedDb, err := gorm.Open(postgres.Open(dns), gormConfig)

		errorsUtils.HandleErrorByPanic(err)

		db = returnedDb

	})

	return db
}

func MigrateDatabase() {
	db := CreateConnection()
	err := db.AutoMigrate(new(models.User), new(models.Todo))

	errorsUtils.HandleErrorByPanic(err)
}
