package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](writer *http.ResponseWriter, r *http.Request) (T, error) {
	var model T
	var defaultValue T
	w := *writer
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		UnprocessableContent(&w, "Validation error")
		return defaultValue, err
	}

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		UnprocessableContent(&w, fmt.Sprintf("Validation error: %s", errors))
		return defaultValue, err
	}
	return model, nil
}
