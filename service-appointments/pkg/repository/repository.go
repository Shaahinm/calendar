package repository

import (
	"github.com/Shaahinm/calendar/config"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any]() *Repository[T] {
	db := config.GetDB()

	return &Repository[T]{
		db: db,
	}
}

func (r *Repository[T]) Create(model T) {
	r.db.Create(&model)
}

func (r *Repository[T]) GetDb() *gorm.DB {
	return r.db
}

func (r *Repository[T]) FindById(id uint) T {
	var entity T
	r.db.First(&entity, id)
	return entity
}

func (r *Repository[T]) FindOneBy(expression string, parameters ...string) T {
	var entity T
	r.db.First(&entity, expression, parameters)
	return entity
}
