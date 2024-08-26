package db

import (
	"github.com/Shaahinm/calendar/config"
	"github.com/Shaahinm/calendar/internal/db/models"
	"github.com/Shaahinm/calendar/pkg/service"
	"gorm.io/gorm"
	// "github.com/Shaahinm/calendar/pkg/service"
)

// alias internal "github.com/Shaahinm/calendar/internal/db/model"

func Up(seed bool) {
	db := config.GetDB()
	db.AutoMigrate(&models.Todo{})

	if seed {
		seedData(db)
	}
}

func seedData(db *gorm.DB) {
	todoService := service.NewService[models.Todo]()
	todoService.Create(models.Todo{Title: "First todo", Description: "description"})
}
