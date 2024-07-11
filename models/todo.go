package models

import (
	"time"
	"todolist/lib/gormLib"
)

type Todo struct {
	Model gormLib.Model
	Title       *string     `json:"title"`
	Done        *bool      `json:"done" gorm:"default:false"`
	Description *string    `json:"description"`
	Deadline    *time.Time `json:"deadline"`
	UserID      uint       `json:"user_id"`
	User        *User      `json:"user,omitempty"`
}
