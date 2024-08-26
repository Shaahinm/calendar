package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Shaahinm/calendar/api/middlewares"
	"github.com/Shaahinm/calendar/api/routes"
	"github.com/Shaahinm/calendar/config"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}

func (s *ApiServer) Start() error {
	router := mux.NewRouter()
	routes.RegisterTodoRoutes(router)
	middlewares.RegisterMiddleware(router)

	connection := fmt.Sprintf("%s:%s", config.Envs.ServerName, config.Envs.Port)
	srv := &http.Server{
		Handler:      router,
		Addr:         connection,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv.ListenAndServe()
}

// func registerMiddleware(router *mux.Router) {
// 	router.Use(mi .loggingMiddleware)
// }
