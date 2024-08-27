package main

import (
	"fmt"
	"log"

	"github.com/Shaahinm/calendar/api"
	"github.com/Shaahinm/calendar/config"
	"github.com/Shaahinm/calendar/internal/db"
)

func main() {
	db.Up(true)
	//InitOtl()
	startServer()
	fmt.Println("Application started")
}

func startServer() {
	baseUrl := fmt.Sprintf("%s:%s", config.Envs.ServerName, config.Envs.Port)
	server := api.NewApiServer(baseUrl)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
