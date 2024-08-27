package routes

import (
	"encoding/json"
	"net/http"

	controllers "github.com/Shaahinm/calendar/api/controllers/todo"
	middlewares "github.com/Shaahinm/calendar/api/middlewares"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/auth/token", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")
	// guarded routes
	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.Use(middlewares.JwtMiddleware)
	authRouter.HandleFunc("/todos/{category:[0-9]+}", controllers.HandleGetTodos).Methods("GET")
	authRouter.HandleFunc("/todos", controllers.HandleCreateTodo).Methods("POST")
}
