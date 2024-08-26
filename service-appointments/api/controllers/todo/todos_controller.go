package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
)

func HandleGetTodos(w http.ResponseWriter, r *http.Request) {
	tracer := otel.GetTracerProvider().Tracer("example")
	_, span := tracer.Start(r.Context(), "my-operation")
	defer span.End()

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
	fmt.Println(vars)
}
