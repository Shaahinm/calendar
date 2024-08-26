package routes

import (
	controllers "github.com/Shaahinm/calendar/api/controllers/todo"
	"github.com/gorilla/mux"
)

func RegisterTodoRoutes(r *mux.Router) {
	r.HandleFunc("/todos/{category:[0-9]+}", controllers.HandleGetTodos).Methods("GET")
	r.HandleFunc("/todos", controllers.HandleCreateTodo).Methods("POST")
}
