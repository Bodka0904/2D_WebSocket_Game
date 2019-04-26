package module

import (
	"log"
)

//Hub holds all Clients
type Hub struct {
	Clients map[*WsClient]bool
}

//Init inits Hub client map
func (h *Hub) Init() {
	h.Clients = make(map[*WsClient]bool)
}

var Hubb = &Hub{}

// GetClients get array of all clients
func (h *Hub) GetClients() []*WsClient {
	keys := make([]*WsClient, 0, len(h.Clients))

	for k := range h.Clients {
		keys = append(keys, k)

	}
	return keys
}

// GetPlayers get array of all players
func (h *Hub) GetPlayers() []*Player {
	keys := make([]*Player, 0, len(h.Clients))

	for k := range h.Clients {
		keys = append(keys, &k.Player)

	}
	return keys
}

// RegisterClient ...
func (h *Hub) RegisterClient(c *WsClient) {
	h.Clients[c] = true
	log.Println("New client registered ", c.Connection.RemoteAddr())
}

// UnregisterClient ...
func (h *Hub) UnregisterClient(c *WsClient) {
	_, ok := h.Clients[c]
	if ok == true {
		delete(h.Clients, c)
		log.Println("Client unregistered ", c.Connection.RemoteAddr())
	}
	return
}
