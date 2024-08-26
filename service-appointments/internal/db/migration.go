package db

import (
	"github.com/Shaahinm/calendar/internal/db/models"
	// "github.com/Shaahinm/calendar/pkg/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// alias internal "github.com/Shaahinm/calendar/internal/db/model"

func Up() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Todo{})

	// test the service
	// todoService := service.NewService[models.Todo]()
	// todoService.Create(models.Todo{Title: "First todo", Description: "description"})
}
