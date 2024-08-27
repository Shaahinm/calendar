package models

type User struct {
	BaseEntity
	Username string
	Password string
}
