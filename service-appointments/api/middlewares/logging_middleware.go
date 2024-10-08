package middlewares

import (
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		// log.Println("---------------------------")
		// log.Println("from middlewares")
		// log.Printf("Method: %s URL: %s", r.Method, r.RequestURI)
		// log.Println("---------------------------")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
