package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shaahinm/calendar/api/routes"
	"github.com/Shaahinm/calendar/config"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Init() {
	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, OpenTelemetry!")
	}), "/"))
	r := mux.NewRouter()
	routes.RegisterTodoRoutes(r)
	test := config.Envs.ServerName
	println(test)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Envs.ServerName + ":" + config.Envs.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
