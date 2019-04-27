package main

import (
	"log"

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
	err = db.CreateInventoryTable(db.Database)
	if err != nil {
		log.Println("Can not create InventoryTable: ", err)
	}
	err = db.AddToInventory(db.Database, "123", "")
	if err != nil {
		log.Println("Can not Add To Inventory: ", err)
	}

	//db.DeleteDbTable(db.Database)

	module.Hubb.Init()
	module.StartAPI()

}
