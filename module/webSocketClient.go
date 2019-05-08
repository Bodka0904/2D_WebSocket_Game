package module

import (
	"log"
	"net/http"
	"time"

	"github.com/gido/2D_WebSocket_Game/db"

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
	Init       bool `default:false`
}

//SendData Every Client sends data about every registered player
func (wsClient *WsClient) SendData() {

	for {
		if !wsClient.Init {
			wsClient.Connection.WriteJSON(WorldList)
			wsClient.Init = true
		} else {
			players := Hubb.GetPlayersInWorld(wsClient.Player.World.Name) //Stores memory addresses of our players
			time.Sleep(30 * time.Millisecond)

			err := wsClient.Connection.WriteJSON(players)
			if err != nil {
				Hubb.UnregisterClient(wsClient)
				wsClient.Connection.Close()
				return
			}
		}
	}
}

//GetData Every Client get data about his player
func (wsClient *WsClient) GetData() {

	for {
		// Reading Commands for movement
		time.Sleep(30 * time.Millisecond)
		err := wsClient.Connection.ReadJSON(&wsClient.Player.Control)
		if err != nil {
			Hubb.UnregisterClient(wsClient)
			wsClient.Connection.Close()
			return
		} else {

			wsClient.Player.UpdatePlayer()
		}

	}
}

//ServeWs It creates and stores clients after socket connection is made
func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)

	}

	wsClient := &WsClient{Connection: conn, Player: Player{ID: "", HP: 20, Energy: 100, Position: Position{}, Velocity: Velocity{X: 3, Y: 3}, World: WorldList[0]}}

	//Get Init message for client Player
	err = wsClient.Connection.ReadJSON(&wsClient.Player)
	if err != nil {
		log.Println(err)
	}

	// Load Inventory of Player
	inv, err := db.GetInventory(db.Database, wsClient.Player.ID)
	if err != nil {
		log.Println(err)
	}

	for _, v := range inv {
		for _, c := range ItemList {
			if v == c.Name {
				// Connect stored names of items in inventory with items from config file
				wsClient.Player.Inventory = append(wsClient.Player.Inventory, Item{v, c.Attack, c.Intellect, c.Defense, c.Level, c.Position})

			}
		}
	}

	//Register Client and his player
	Hubb.RegisterClient(wsClient)

	//Send data about all players
	go wsClient.SendData()

	//Recieve data from client about particular player
	go wsClient.GetData()
}
