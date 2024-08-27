package controllers

import (
	"net/http"

	"github.com/Shaahinm/calendar/api/controllers"
	"github.com/Shaahinm/calendar/internal/db/models"
	"github.com/Shaahinm/calendar/pkg/service"

	"go.opentelemetry.io/otel"
)

func HandleGetTodos(w http.ResponseWriter, r *http.Request) {
	tracer := otel.GetTracerProvider().Tracer("service-appointments")
	_, span := tracer.Start(r.Context(), "getTodos")
	defer span.End()
	println(tracer)
	//vars := mux.Vars(r)
	controllers.Ok(&w, "todos")
}

func HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	model, err := controllers.Validate[models.Todo](&w, r)
	if err == nil {
		todoService := service.NewTodoService()
		t := todoService.Create(model)
		controllers.Ok(&w, t)
	}
}
