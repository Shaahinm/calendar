package pkg

import (
	repository "github.com/Shaahinm/calendar/pkg/repository"
)

func NewService[T any]() *Service[T] {
	repository := *repository.NewRepository[T]()
	return &Service[T]{repository: &repository}
}

type Service[T any] struct {
	repository *repository.Repository[T]
}

func (s *Service[T]) Create(model T) {
	s.repository.Create(model)
}
