package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shaahinm/calendar/api/controllers"
	"github.com/Shaahinm/calendar/api/controllers/todo"
	"github.com/Shaahinm/calendar/pkg/util"

	middlewares "github.com/Shaahinm/calendar/api/middlewares"
	"github.com/gorilla/mux"
)

type Test struct {
	Id      int
	Message string
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/auth/token", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		url := "http://localhost:8000/test"
		result := util.Get[[]Test](url)
		fmt.Println(result)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler

		t := make([]Test, 0)
		t = append(t, Test{1, "hello"})
		t = append(t, Test{2, "world"})

		controllers.Ok(&w, t)
	}).Methods("GET")
	// guarded routes
	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.Use(middlewares.JwtMiddleware)
	authRouter.HandleFunc("/todos/{category:[0-9]+}", todo.HandleGetTodos).Methods("GET")
	authRouter.HandleFunc("/todos", todo.HandleCreateTodo).Methods("POST")
}
