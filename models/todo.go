package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title       string
	Done        *bool
	Description *string
	Deadline    *time.Time
	UserID      uint
}
