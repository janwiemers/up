package helper

import (
	"time"

	"github.com/janwiemers/up/database"
)

// Cleanup takes care of cleaning old checks in a go routine
func Cleanup() {
	database.CleanupChecks()
	time.Sleep(1 * time.Hour)
	go Cleanup()
}
