package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shaahinm/calendar/api/routes"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Init() {
	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, OpenTelemetry!")
	}), "/"))
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
