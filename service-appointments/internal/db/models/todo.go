package models

type Todo struct {
	BaseEntity
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description"`
}
