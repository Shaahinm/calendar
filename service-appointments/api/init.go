package api

import (
	"log"
	"net/http"
	"time"

	todo "github.com/Shaahinm/calendar/api/controllers/todo"
	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/todos/{category}", todo.HandleGetTodos).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
