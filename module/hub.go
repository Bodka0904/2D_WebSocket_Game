package module

import (
	"log"
	"sort"
)

//Hub holds all Clients
type Hub struct {
	Clients map[*WsClient]int
}

//Init inits Hub client map
func (h *Hub) Init() {
	h.Clients = make(map[*WsClient]int)
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
	hack := map[int]*Player{}

	hackkeys := []int{}

	for key, val := range h.Clients {

		hack[val] = &key.Player
		hackkeys = append(hackkeys, val)

	}
	sort.Ints(hackkeys)

	for _, val := range hackkeys {

		keys = append(keys, hack[val])
	}

	return keys
}

func (h *Hub) GetPlayersInWorld(WorldName string) []*Player {
	keys := make([]*Player, 0, len(h.Clients))
	hack := map[int]*Player{}

	hackkeys := []int{}

	for key, val := range h.Clients {
		if key.Player.World.Name == WorldName {

			hack[val] = &key.Player
			hackkeys = append(hackkeys, val)
		}
	}
	sort.Ints(hackkeys)

	for _, val := range hackkeys {

		keys = append(keys, hack[val])
	}

	return keys
}

// RegisterClient ...
func (h *Hub) RegisterClient(c *WsClient) {
	h.Clients[c] = len(h.Clients)

	log.Println("New client registered ", c.Connection.RemoteAddr())
}

// UnregisterClient ...
func (h *Hub) UnregisterClient(c *WsClient) {

	delete(h.Clients, c)
	log.Println("Client unregistered ", c.Connection.RemoteAddr())

	return
}
