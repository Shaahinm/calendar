package middlewares

import "github.com/gorilla/mux"

func RegisterMiddleware(router *mux.Router) {
	router.Use(loggingMiddleware)
}
