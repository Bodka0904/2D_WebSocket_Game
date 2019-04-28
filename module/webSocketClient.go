package module

import (
	"log"
	"net/http"
	"time"

	"github.com/gido/2D_WebSocket_Game/server/db"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//WsClient ws connection and player
type WsClient struct {
	Connection *websocket.Conn
	Player     Player
}

//SendData Every Client sends data about every registered player
func (wsClient *WsClient) SendData() {

	for {

		players := Hubb.GetPlayers() //Stores memory addresses of our players
		time.Sleep(25 * time.Millisecond)

		err := wsClient.Connection.WriteJSON(players)
		if err != nil {
			Hubb.UnregisterClient(wsClient)
			wsClient.Connection.Close()
			return
		}

	}
}

//GetData Every Client get data about his player
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

//ServeWs It creates and stores clients after socket connection is made
func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)

	}

	wsClient := &WsClient{Connection: conn, Player: Player{ID: "", Position: Position{}, Velocity: Velocity{X: 3, Y: 3}}}

	//Get Init message for client Player
	err = wsClient.Connection.ReadJSON(&wsClient.Player)
	if err != nil {
		log.Println(err)
	}
	err, _ = db.GetInventory(db.Database, wsClient.Player.ID)
	if err != nil {
		log.Println(err)
	}

	//Register Client and his player
	Hubb.RegisterClient(wsClient)

	//Send data about all players
	go wsClient.SendData()

	//Recieve data from client about particular player
	go wsClient.GetData()
}
