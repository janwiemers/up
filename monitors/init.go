package monitors

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/janwiemers/up/database"
	"github.com/janwiemers/up/models"
	"github.com/janwiemers/up/notifications"
	"github.com/janwiemers/up/websockets"
	"github.com/spf13/viper"
)

// InitAllMonitors initializes the Monitors for all applications given in the `monitors` argument
func InitAllMonitors(monitors []models.Application) {
	for monitor := range monitors {
		app := populateApplicationDefaults(monitors[monitor])

		// create application in DB
		app = database.CreateAndUpdateApplication(app)

		// start monitor
		go InitMonitor(app.ID, 0)
	}
}

// InitMonitor initializes a singe monitor
func InitMonitor(id int, retry int) {
	var up bool
	app := database.GetApplication(id)
	if app.Protocol == "http" {
		statusCode, err := strconv.Atoi(app.Expectation)
		if err != nil {
			log.Fatalln(fmt.Sprintf("%v: Cannot start monitor", app.Name))
		}
		up = httpMonitor(app.Target, statusCode)
	}

	if app.Protocol == "dns" {
		up = dnsMonitor(app.Target, app.Expectation)
	}

	if up != true {
		log.Println(fmt.Sprintf("%v: %v", app.Name, up))
		if retry == viper.GetInt("MAX_RETRY") && app.Degraded == false {
			go InitMonitor(app.ID, 0)
			app = database.ApplicationSetDegraded(app, true)
			notifications.Dispatch(app.Name, true)
			return
		}
		time.Sleep(10 * time.Second)
		c, _ := database.InsertCheck(app, up)
		websockets.BroadcastCheck(*c)
		go InitMonitor(app.ID, retry+1)
		return
	}

	time.Sleep(app.Interval)
	if app.Degraded == true {
		notifications.Dispatch(app.Name, false)
		app = database.ApplicationSetDegraded(app, false)
	}

	c, _ := database.InsertCheck(app, up)
	websockets.BroadcastCheck(*c)
	go InitMonitor(app.ID, 0)
}

func populateApplicationDefaults(application models.Application) models.Application {
	if application.Protocol == "" {
		application.Protocol = "http"
	}

	if application.Expectation == "" {
		application.Expectation = "200"
	}

	if application.Interval == 0 {
		application.Interval = 5 * time.Minute
	}

	return application
}

func sendMessage(t string, payload []byte) {
	j, _ := json.Marshal(fmt.Sprintf("{type: \"%v\", message: %v}", t, payload))
	websockets.HubInstance.Broadcast <- j
}
