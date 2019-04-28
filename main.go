package main

import (
	"log"

	"github.com/gido/2D_WebSocket_Game/server/config"
	"github.com/gido/2D_WebSocket_Game/server/db"
	"github.com/gido/2D_WebSocket_Game/server/module"
)

func main() {

	err := db.InitDB()
	if err != nil {
		log.Println("Can not Init Database: ", err)
	}

	err = db.CreateDbTable(db.Database)
	if err != nil {
		log.Println("Can not Init Database: ", err)
	}
	err = db.CreateInventoryTable(db.Database, "cMVrjkJWXL")
	if err != nil {
		log.Println("Can not create InventoryTable: ", err)
	}
	err = db.AddToInventory(db.Database, "cMVrjkJWXL", "tesfght")
	if err != nil {
		log.Println("Can not Add To Inventory: ", err)
	}
	var test []string
	for _, v := range module.Hubb.GetPlayers() {
		test = append(test, v.ID)

	}
	test = append(test, "drn7dnvjgt")

	//db.DeleteDbTable(db.Database, test)
	config.Init()
	module.Hubb.Init()
	module.StartAPI()

}
