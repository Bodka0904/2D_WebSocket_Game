package main

import (
	"fmt"

	"github.com/gido/2D_WebSocket_Game/server/db"
	"github.com/gido/2D_WebSocket_Game/server/module"
)

func main() {

	err := db.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	err = db.CreateDbTable(db.Database)
	if err != nil {
		fmt.Println(err)
	}
	err = db.CreateInventoryTable(db.Database)
	if err != nil {
		fmt.Println(err)
	}
	err = db.AddNewInventorySlot(db.Database, "123")
	if err != nil {
		fmt.Println(err)
	}
	//db.LoginPlayer(db.Database, "test", "test")
	//db.DeleteDbTable(db.Database)

	module.Hubb.Init()
	module.StartApi()

}
