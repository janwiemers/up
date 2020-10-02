package database

import (
	"time"

	"github.com/janwiemers/up/models"
	"github.com/spf13/viper"
)

// Checks retrieves all applications from the database
func Checks(id int) ([]models.Check, error) {
	db := connect()

	var checks []models.Check
	err := db.Order("id desc").Limit(24).Select("id", "up", "created_at", "application_id").Find(&checks, "application_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return checks, nil
}

// CleanupChecks deletes all checks that are older then the threshold defined in `DB_CLEANUP_AFTER`
func CleanupChecks() bool {
	db := connect()

	var check models.Check
	t := time.Now().Sub(viper.GetTime("DB_CLEANUP_AFTER"))
	err := db.Where("created_at <= ?", t).Delete(&check).Error
	if err != nil {
		return false
	}
	return true
}

// DeleteChecks deletes all checks from an application
func DeleteChecks(id int) error {
	db := connect()
	var check models.Check

	err := db.Delete(&check, "application_id", id).Error
	if err != nil {
		return err
	}
	return nil
}
