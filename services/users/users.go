package usersService

import (
	"booking/constants"
	"booking/lib/gormLib"
	"booking/models"
	"booking/utils"
	"errors"
	"gorm.io/gorm"
)

func FindUserByUsername(username string) (user models.User, err error) {
	db := gormLib.CreateConnection()

	result := db.Where(&models.User{Username: username}).First(&user)

	err = result.Error

	return user, err
}

func Login(username string, password string) (user models.User, err error) {
	user, err = FindUserByUsername(username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, constants.ErrUserNotFound
		}
		return user, err
	}

	if !utils.VerifyHashedText(user.Password, password) {
		return user, constants.ErrInvalidCredentials
	}

	return user, err
}

func Signup(username string, password string) (user models.User, err error) {
	db := gormLib.CreateConnection()

	user, err = FindUserByUsername(username)

	// If FindUserByUsername didn't return an error it means that it found a user which it shouldn't in sing up
	//And if the error is the record not found it is a good news that user doesn't exist. otherwise it is an important error
	if err == nil {
		return user, constants.ErrUserAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, err
	}

	user = models.User{Username: username, Password: password}

	result := db.Create(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil

}
