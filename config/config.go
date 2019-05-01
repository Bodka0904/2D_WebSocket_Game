package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//ConfigItem stores information about all Items loaded from Item.json file
type ConfigItem struct {
	Name      string
	Attack    int
	Intellect int
	Defense   int
}

var Items []ConfigItem

//LoadItems loads Items.json file with information about all items in game
func LoadItems() error {
	file, err := ioutil.ReadFile("Items.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &Items)
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
