package monitors

import (
	"github.com/janwiemers/up/database"
	"github.com/janwiemers/up/models"
)

// Cleanup deletes monitors that are not in the config file
func Cleanup(monitors []models.Application) {
	apps := database.Applications()
	for i := range apps {
		app := apps[i]

		if contains(app.Name, monitors) == false {
			database.DeleteApplication(app.ID)
			database.DeleteChecks(app.ID)
		}
	}
}

func contains(name string, monitors []models.Application) bool {
	for m := range monitors {
		monitor := monitors[m]
		if monitor.Name == name {
			return true
		}
	}
	return false
}
