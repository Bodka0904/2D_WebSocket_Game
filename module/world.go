package module

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

const (
	Width  = 1000
	Height = 800
)

type World struct {
	Name      string
	Level     int
	Creatures []Creature
	Resources []Resource
	Players   []Player

	//Add copies of players that are in same world as client keep track of them and send world via wsClient.Player.World with creatures and Resources in it
}

type Resource struct {
	Name     string
	Capacity int
	Rare     int
	Respawn  int
	Position Position
}

type Creature struct {
	Name       string
	Level      int
	Friend     bool
	HP         int
	Position   Position
	Velocity   Velocity
	Attributes Attributes
	XP         int
	Inventory  []Item
	Respawn    int //Seconds
}

//ResourceList list of all avaiable Resources
var ResourceList []Resource

// CreatureList list of all avaiable Creatures
var CreatureList []Creature

// ItemList list of all avaiable Items
var ItemList []Item

// WorldList list of all avaiable Worlds
var WorldList []*World

//Init init world configuration
func Init() {
	err := LoadItems()
	if err != nil {
		log.Println(err)
	}
	err = LoadCreatures()
	if err != nil {
		log.Println(err)
	}
	err = LoadWorlds()
	if err != nil {
		log.Println(err)
	}
	err = LoadResources()
	if err != nil {
		log.Println(err)
	}
	for _, v := range WorldList {
		v.GenerateCreatures()
		v.GenerateResources()
	}
	for _, v := range CreatureList {
		v.GenerateItems()
	}

}

func LoadItems() error {
	file, err := ioutil.ReadFile("Items.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &ItemList)
	if err != nil {
		return err
	}

	return nil
}

func (c *Creature) GenerateItems() {

	rand.Seed(time.Now().UnixNano())
	num := 3
	rangeItems := rand.Intn(num)

	for i := 0; i < rangeItems; i++ {
		rand.Seed(time.Now().UnixNano())
		num = len(ItemList)
		randomNum := rand.Intn(num)

		c.Inventory = append(c.Inventory, ItemList[randomNum])
	}

}

func (c *Creature) DropItems() {
	for _, v := range c.Inventory {
		v.Position = c.Position
	}
}

func LoadCreatures() error {
	file, err := ioutil.ReadFile("Creatures.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &CreatureList)
	if err != nil {
		return err
	}

	return nil
}

func (w *World) GenerateCreatures() {

	for _, v := range CreatureList {
		if v.Level <= w.Level {

			w.Creatures = append(w.Creatures, v)

		}
	}
}
func (w *World) GenerateResources() {

	rand.Seed(time.Now().UnixNano())
	min := 50
	max := 750

	for _, v := range ResourceList {
		for i := 0; i < v.Rare; i++ {

			v.Position.X = float64(rand.Intn(max-min) + min)
			v.Position.Y = float64(rand.Intn(max-min) + min)
			w.Resources = append(w.Resources, v)
		}
	}
}

func LoadResources() error {
	file, err := ioutil.ReadFile("Resources.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &ResourceList)
	if err != nil {
		return err
	}

	return nil
}

func LoadWorlds() error {
	file, err := ioutil.ReadFile("Worlds.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &WorldList)
	if err != nil {
		return err
	}

	return nil
}
