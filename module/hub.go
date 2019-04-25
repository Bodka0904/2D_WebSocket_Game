package module

import "fmt"

type Hub struct {
	Clients map[*WsClient]bool
}

func (h *Hub) Init() {
	h.Clients = make(map[*WsClient]bool)

}

var Hubb = &Hub{}

func (h *Hub) getClients() []*WsClient {
	keys := make([]*WsClient, 0, len(h.Clients))

	for k := range h.Clients {
		keys = append(keys, k)

	}
	return keys
}

func (h *Hub) getPlayers() []*Player {
	keys := make([]*Player, 0, len(h.Clients))

	for k := range h.Clients {
		keys = append(keys, &k.Player)

	}
	return keys
}

func (h *Hub) RegisterClient(c *WsClient) {
	h.Clients[c] = true
	fmt.Println("New client registered", c.Connection.RemoteAddr())
}

func (h *Hub) UnregisterClient(c *WsClient) {
	_, ok := h.Clients[c]
	if ok == true {
		delete(h.Clients, c)
		fmt.Println("Client unregistered", c.Connection.RemoteAddr())
	}
	return
}
