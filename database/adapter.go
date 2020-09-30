package database

import (
	"github.com/janwiemers/up/models"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect establishes a connection to the Database
func connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(viper.GetString("DB_PATH")), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Application{})
	db.AutoMigrate(&models.Check{})

	return db
}

// CreateAndUpdateApplication creates an application record if it does not exist already
func CreateAndUpdateApplication(application models.Application) models.Application {
	db := connect()
	var app models.Application
	db.FirstOrCreate(&app, "name = ?", application.Name)
	db.Model(&app).Updates(application)
	return app
}

// InsertCheck inserts a new check into the databse
func InsertCheck(application models.Application, up bool) (bool, error) {
	db := connect()
	var app models.Application
	db.Select([]string{"ID", "Name"}).First(&app, "name = ?", application.Name)
	db.Create(&models.Check{UP: up, ApplicationID: app.ID})

	return true, nil
}
