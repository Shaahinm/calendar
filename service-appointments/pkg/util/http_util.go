package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Get[T any](url string) T {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

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
