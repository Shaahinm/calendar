package main

import (
	"fmt"
	"log"

	"github.com/Shaahinm/calendar/api"
	"github.com/Shaahinm/calendar/config"
	"github.com/Shaahinm/calendar/internal/db"
	lang "github.com/Shaahinm/calendar/internal/locale"
)

func main() {
	db.Up(false)
	// InitOtl()
	startServer()
	fmt.Println("Application started")
}

func startServer() {
	// lang.Init()
	log.Println(lang.Get(lang.HelloWorld))
	template := map[string]interface{}{
		"Name": "Shaahin",
	}
	log.Println(lang.GetWithTemplate(lang.Test, template))
	baseUrl := fmt.Sprintf("%s:%s", config.Envs.ServerName, config.Envs.Port)
	server := api.NewApiServer(baseUrl)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
