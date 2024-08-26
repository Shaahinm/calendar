package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shaahinm/calendar/internal/db/models"
	"github.com/Shaahinm/calendar/pkg/service"
	"github.com/go-playground/validator/v10"

	"go.opentelemetry.io/otel"
)

func HandleGetTodos(w http.ResponseWriter, r *http.Request) {
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
	//

	result := Result{&w}
	result.ok("Hello")
}

func HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(todo)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Todo created successfully!")
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
