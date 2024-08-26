package main

import (
	"fmt"

	"github.com/Shaahinm/calendar/api"
	internal "github.com/Shaahinm/calendar/internal/db"
)

func main() {
	internal.Up()
	api.Init()
	fmt.Println("Application started")
}
