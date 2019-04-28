package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gido/2D_WebSocket_Game/server/module"
)

//ListItems stores information about all Items loaded from Item.json file
var ListItems []module.Item

//LoadItems loads Items.json file with information about all items in game
func LoadItems() error {
	file, err := ioutil.ReadFile("Items.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &ListItems)
	if err != nil {
		return err
	}

	return nil
}

//Init init configuration
func Init() {
	err := LoadItems()
	if err != nil {
		log.Println(err)
	}
}
