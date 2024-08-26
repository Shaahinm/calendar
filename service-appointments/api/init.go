package api

import (
	"log"
	"net/http"
	"time"

	"github.com/Shaahinm/calendar/api/routes"
	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()
	routes.RegisterTodoRoutes(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
