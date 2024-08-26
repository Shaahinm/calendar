package main

import (
	"fmt"
	"log"

	"github.com/Shaahinm/calendar/api"
	"github.com/Shaahinm/calendar/config"
	"github.com/Shaahinm/calendar/internal/db"
)

func main() {
	config.ConnectToDatabase()
	database := config.GetDB()

	db.Up(false)
	InitOtl()

	connectionString := fmt.Sprintf("%s:%s", config.Envs.ServerName, config.Envs.Port)
	server := api.NewApiServer(connectionString, database)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	// db.Up()
	// api.Init()
	// InitOtl()
	fmt.Println("Application started")
}
