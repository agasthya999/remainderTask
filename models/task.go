package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"primary_key"`
	Description string
	Priority    int
	DueDate     int //epoch time
	UserID      int
	ContactId   int
}
