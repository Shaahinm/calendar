package service

import (
	"github.com/Shaahinm/calendar/pkg/repository"
)

func NewService[T any]() *Service[T] {
	repository := *repository.NewRepository[T]()
	return &Service[T]{&repository}
}

type Service[T any] struct {
	//repository *repository.Repository[T]
	*repository.Repository[T]
}

// func (s *Service[T]) Create(model T) T {
// 	t := s.repository.Create(model)
// 	fmt.Println(t)
// 	return t
// }
