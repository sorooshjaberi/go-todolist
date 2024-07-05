package gormLib

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"todolist/constants"
	"todolist/lib/dotenvLib"
	"todolist/models"
	"todolist/utils/errorsUtils"
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

func GetOffsetUsingPageSize(page int, perPage int) int {
	return (page - 1) * perPage
}

func Paginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		offset := GetOffsetUsingPageSize(page, perPage)
		return db.Offset(offset).Limit(perPage)
	}
}
