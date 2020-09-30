package monitors

import (
	"github.com/janwiemers/up/database"
	"github.com/janwiemers/up/models"
)

// Cleanup deletes monitors that are not in the config file
func Cleanup(monitors models.Monitors) {
	apps := database.Applications()
	for i := range apps {
		app := apps[i]

		if contains(app.Name, monitors) == false {
			database.DeleteApplication(app.ID)
			database.DeleteChecks(app.ID)
		}
	}
}

func contains(name string, monitors models.Monitors) bool {
	for m := range monitors.Applications {
		monitor := monitors.Applications[m]
		if monitor.Name == name {
			return true
		}
	}
	return false
}
