package service

import "github.com/Shaahinm/calendar/internal/db/models"

func NewTodoService() *TodoService {
	service := *NewService[models.Todo]()
	return &TodoService{&service}
}

type TodoService struct {
	*Service[models.Todo]
}

func (r *TodoService) custom(result any) {

}
