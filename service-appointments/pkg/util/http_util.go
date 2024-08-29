package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Todo: to make the StandardContainer in response.go to use this one instead (generic)
type StandardContainer[T any] struct {
	Data    T   `json:"data"`
	Code    int `json:"code"`
	Version int `json:"version"`
}

func Get[T any](url string) T {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var t StandardContainer[T]
	json.Unmarshal(body, &t)
	return t.Data
}

func GetPlain[T any](url string) T {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var t T
	json.Unmarshal(body, &t)
	return t
}

func Post[T any](result any) {

}
