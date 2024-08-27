package models

type Todo struct {
	BaseModel
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description"`
}
