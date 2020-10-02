package websockets

import (
	"encoding/json"

	"github.com/janwiemers/up/models"
)

type BroadcastData struct {
	Type     string               `json:"type"`
	Monitors []models.Application `json:"monitors"`
	Check    models.Check         `json:"check"`
	Checks   []models.Check       `json:"checks"`
}

func NewBroadcastDataForMonitor(t string, monitors []models.Application) *BroadcastData {
	return &BroadcastData{
		Type:     "monitors",
		Monitors: monitors,
	}
}

// BroadcastMonitors publishes all monitors to all connected websocket clients
func BroadcastMonitors(monitors []models.Application) {
	var payload = &BroadcastData{
		Type:     "monitors",
		Monitors: monitors,
	}
	j, _ := json.Marshal(payload)
	HubInstance.Broadcast <- j
}

// BroadcastCheck sends a new check update to all connected websocket clients
func BroadcastCheck(check models.Check) {
	var payload = &BroadcastData{
		Type:  "addCheck",
		Check: check,
	}
	j, _ := json.Marshal(payload)
	HubInstance.Broadcast <- j
}
