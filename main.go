package main

import (
	"github.com/gido/2D_WebSocket_Game/db"
	"github.com/gido/2D_WebSocket_Game/module"
)

func main() {
	db.InitDB()
	//db.DeleteDbSchema(db.Database)

	module.Init()
	module.Hubb.Init()
	module.StartAPI()

}
