package internal

import (
	internal "github.com/Shaahinm/calendar/internal/db/model"
	pkg "github.com/Shaahinm/calendar/pkg/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// alias internal "github.com/Shaahinm/calendar/internal/db/model"

func Up() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&internal.Todo{})

	// test the service
	todoService := pkg.NewService[internal.Todo]()
	todoService.Create(internal.Todo{Title: "First todo", Description: "description"})
}
