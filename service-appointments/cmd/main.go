package main

import (
	"fmt"

	"github.com/Shaahinm/calendar/api"
	"github.com/Shaahinm/calendar/internal/db"
)

func main() {
	db.Up()
	api.Init()
	InitOtl()
	fmt.Println("Application started")
}
