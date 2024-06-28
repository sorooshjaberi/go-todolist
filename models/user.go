package models

import (
	"booking/utils/encryptionUtil"
	"booking/utils/errorsUtils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := encryptionUtil.HashText(u.Password)
		errorsUtils.HandleErrorByPanic(err)
		u.Password = hashedPassword
	}
	return nil
}
