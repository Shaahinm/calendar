package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Shaahinm/calendar/internal/db/models"
	"github.com/Shaahinm/calendar/pkg/service"
	"go.opentelemetry.io/otel"
)

func HandleGetTodos(w http.ResponseWriter, r *http.Request) {
	// TODO: to add to middleware
	// input validation
	tracer := otel.GetTracerProvider().Tracer("service-appointments")
	_, span := tracer.Start(r.Context(), "getTodos")
	defer span.End()
	println(tracer)

	// vars := mux.Vars(r)

	// todoService := service.NewService[models.Todo]()
	// todo := models.Todo{Title: vars["category"], Description: "description"}
	// t := todoService.Repository().Create(todo)

	// result := Result{&w}
	// result.ok(t)
}

type Result struct {
	writer *http.ResponseWriter
}

func (r *Result) ok(result any) {
	writer := *r.writer
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(result)
	writer.Write(res)
}

type TodoService struct {
	todoService service.Service[models.Todo]
}

// func (r *TodoService) custom(result any) {
// 	r.todoService.Repository().GetDb().
// }
