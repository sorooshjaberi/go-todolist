package models

import (
	"booking/constants"
	"booking/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := utils.HashText(u.Password)
		constants.HandleErrorByPanic(err)
		u.Password = hashedPassword
	}
	return nil
}
