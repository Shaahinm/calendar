package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description"`
}
