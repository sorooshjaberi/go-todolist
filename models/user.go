package models

import (
	"booking/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := utils.HashText(u.Password)
		utils.HandleErrorByPanic(err)
		u.Password = hashedPassword
	}
	return nil
}
