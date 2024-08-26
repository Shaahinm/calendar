package routes

import (
	controllers "github.com/Shaahinm/calendar/api/controllers/todo"
	"github.com/gorilla/mux"
)

func RegisterTodoRoutes(r *mux.Router) {
	r.HandleFunc("/todos/{category}", controllers.HandleGetTodos).Methods("GET")
}
