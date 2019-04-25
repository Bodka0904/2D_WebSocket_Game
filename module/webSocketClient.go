package module

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsClient struct {
	Connection *websocket.Conn
	Player     Player
}

func (wsClient *WsClient) SendData() {

	for {

		players := Hubb.getPlayers() //Stores memory addresses of our players
		time.Sleep(25 * time.Millisecond)

		err := wsClient.Connection.WriteJSON(players)
		if err != nil {
			Hubb.UnregisterClient(wsClient)
			wsClient.Connection.Close()
			return
		}

	}
}

func (wsClient *WsClient) GetData() {

	for {
		// Reading Commands for movement
		time.Sleep(25 * time.Millisecond)
		err := wsClient.Connection.ReadJSON(&wsClient.Player.Control)
		if err != nil {
			Hubb.UnregisterClient(wsClient)
			wsClient.Connection.Close()
			return
		} else {
			wsClient.Player.UpdatePosition()
		}

	}
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)

	}

	wsClient := &WsClient{Connection: conn}
	router.HandleFunc("/loginHandler", wsClient.LoginHandler).Methods("POST")

	Hubb.RegisterClient(wsClient)

	go wsClient.SendData()
	go wsClient.GetData()

}
