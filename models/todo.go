package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title       string     `json:"title"`
	Done        *bool      `json:"done"`
	Description *string    `json:"description"`
	Deadline    *time.Time `json:"deadline"`
	UserID      uint       `json:"user_id"`
	User        *User      `json:"user,omitempty"`
}
