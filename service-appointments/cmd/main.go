package main

import (
	"fmt"
	"log"

	"github.com/Shaahinm/calendar/api"
	"github.com/Shaahinm/calendar/config"
	"github.com/Shaahinm/calendar/internal/db"
)

func main() {
	//config.ConnectToDatabase()

	db.Up(false)
	InitOtl()

	baseUrl := fmt.Sprintf("%s:%s", config.Envs.ServerName, config.Envs.Port)
	server := api.NewApiServer(baseUrl)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	// db.Up()
	// api.Init()
	// InitOtl()
	fmt.Println("Application started")
}
