package models

import (
	"gorm.io/gorm"
	"todolist/lib/gormLib"
	"todolist/utils/encryptionUtils"
	"todolist/utils/errorsUtils"
)

type User struct {
	gormLib.Model
	Username string `json:"username" binding:"required" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := encryptionUtils.HashText(u.Password)
		errorsUtils.HandleErrorByPanic(err)
		u.Password = hashedPassword
	}
	return nil
}
