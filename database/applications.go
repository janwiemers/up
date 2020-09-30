package database

import (
	"github.com/janwiemers/up/models"
)

// Applications retrieves all applications from the database
func Applications() []models.Application {
	db := connect()
	var applications []models.Application
	db.Order("name asc").Find(&applications)

	return applications
}

// ApplicationSetDegraded sets an application to the degraded state
func ApplicationSetDegraded(app models.Application, degraded bool) models.Application {
	db := connect()
	err := db.First(&app, "name = ?", app.Name).Update("degraded", degraded).Error

	if err != nil {
		panic("not able to update")
	}
	return app
}

// GetApplication returns an application
func GetApplication(id int) models.Application {
	db := connect()
	var app models.Application
	err := db.First(&app, id).Error
	if err != nil {
		panic("foo")
	}

	return app
}

// DeleteApplication deletes an application
func DeleteApplication(id int) error {
	db := connect()
	var app models.Application

	err := db.Delete(&app, id).Error
	if err != nil {
		return err
	}
	return nil
}
