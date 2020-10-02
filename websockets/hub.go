package websockets

import (
	"encoding/json"

	"github.com/janwiemers/up/database"
)

var HubInstance *Hub

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			apps := database.Applications()
			d, _ := json.Marshal(NewBroadcastDataForMonitor("monitors", apps))
			client.send <- d

			for i := range apps {
				app := apps[i]
				checks, _ := database.Checks(app.ID)
				d, _ := json.Marshal(&BroadcastData{
					Type:   "addChecks",
					Checks: checks,
				})
				client.send <- d
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
