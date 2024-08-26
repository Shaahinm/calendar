package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Shaahinm/calendar/api/routes"
	"github.com/Shaahinm/calendar/config"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ApiServer struct {
	addr string
	db   *gorm.DB
}

func NewApiServer(addr string, db *gorm.DB) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}

func (s *ApiServer) Start() error {
	router := mux.NewRouter()
	routes.RegisterTodoRoutes(router)

	connection := fmt.Sprintf("%s:%s", config.Envs.ServerName, config.Envs.Port)
	srv := &http.Server{
		Handler:      router,
		Addr:         connection,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv.ListenAndServe()
}

// func Init() {
// 	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello, OpenTelemetry!")
// 	}), "/"))
// 	r := mux.NewRouter()
// 	routes.RegisterTodoRoutes(r)
//
// 	srv := &http.Server{
// 		Handler:      r,
// 		Addr:         config.Envs.ServerName + ":" + config.Envs.Port,
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}
//
// 	log.Fatal(srv.ListenAndServe())
// }
