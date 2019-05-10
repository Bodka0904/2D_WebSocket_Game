package module

import (
	"crypto/rand"
	"log"
	"math/big"
)

type Player struct {
	ID              string
	HP              int
	Energy          float64
	Position        Position
	Velocity        Velocity
	Control         Control
	Attributes      Attributes
	Face            string `default:"Right"`
	BuildMode       bool
	BonusAttributes BonusAttributes
	Inventory       []Item
	World           *World
}

func (p *Player) UpdatePlayer() {

	if p.Control.Up {
		p.Position.Y -= p.Velocity.Y
		p.Face = "Up"
	}
	if p.Control.Down {
		p.Position.Y += p.Velocity.Y
		p.Face = "Down"
	}
	if p.Control.Left {
		p.Position.X -= p.Velocity.X
		p.Face = "Left"
	}
	if p.Control.Right {
		p.Position.X += p.Velocity.X
		p.Face = "Right"
	}
	p.ChangeWorld(WorldList[1])

	if p.Control.Action.Attack {
		if p.Energy >= 2 {
			p.Energy -= 2
			p.BuildMode = false

		} else {
			p.Control.Action.Attack = false
		}
	}
	if p.Control.Action.Mine {
		if p.Energy >= 2 {
			p.Energy -= 2
			p.BuildMode = false
		} else {
			p.Control.Action.Mine = false
		}
	}
	if p.Control.Action.Build {
		if p.Energy >= 10 {
			p.Build()
			p.Energy -= 10
		} else {
			p.Control.Action.Build = false
			p.BuildMode = false
		}
	}
	if p.Energy < 100 {
		p.Energy += 0.5
	}

	p.PickItem()

}

// UpdateAttributes calculate bonus attributes of items
func (p *Player) UpdateAttributes() {

	var bonusAttack, bonusIntellect, bonusDefense int

	for _, v := range p.Inventory {
		bonusAttack += v.Attack
		bonusIntellect += v.Intellect
		bonusDefense += v.Defense
	}

	p.BonusAttributes.Attack += bonusAttack
	p.BonusAttributes.Intellect += bonusIntellect
	p.BonusAttributes.Defense += bonusDefense
}

//GetToken generates ID for player
func GetToken(length int) string {
	token := ""
	codeAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	codeAlphabet += "abcdefghijklmnopqrstuvwxyz"
	codeAlphabet += "0123456789"

	for i := 0; i < length; i++ {
		token += string(codeAlphabet[randString(int64(len(codeAlphabet)))])
	}
	return token
}

// randString random string for GetToken function
func randString(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println(err)
	}
	return nBig.Int64()
}

func (p *Player) ChangeWorld(World *World) {
	if p.Position.X < -20 || p.Position.X > Width+20 || p.Position.Y < -20 || p.Position.Y > Height+20 {
		p.World = World
		p.Position.X = Width - p.Position.X
		p.Position.Y = Height - p.Position.Y
	}
}

func (p *Player) PickItem() {
	if p.Control.Action.Pick {
		for v := 0; v < len(p.World.Items); v++ {
			if p.Position.X <= (p.World.Items[v].Position.X+5) && p.World.Items[v].Position.X >= (p.World.Items[v].Position.X-5) && p.World.Items[v].Position.Y <= (p.World.Items[v].Position.Y+5) && p.Position.Y >= (p.World.Items[v].Position.Y-5) {
				p.Inventory = append(p.Inventory, p.World.Items[v])
				p.World.Items = append(p.World.Items[:v], p.World.Items[v+1:]...)
			}
		}
	}
}

func (p *Player) Build() {
	p.Inventory = append(p.Inventory, ItemList[3]) //DELETE

	p.BuildMode = true

	if p.Control.Action.Build {
		if p.Face == "Up" {
			if p.Inventory[p.Control.Action.SelectedItem].Type == "Material" {
				p.World.Objects = append(p.World.Objects, Object{p.Inventory[p.Control.Action.SelectedItem].Name, 10, Position{p.Position.X + 5, p.Position.Y - 15}, p.Inventory[p.Control.Action.SelectedItem].Width, p.Inventory[p.Control.Action.SelectedItem].Height})
				p.Inventory = append(p.Inventory[:p.Control.Action.SelectedItem], p.Inventory[p.Control.Action.SelectedItem+1:]...)
			}
		}
		if p.Face == "Down" {
			if p.Inventory[p.Control.Action.SelectedItem].Type == "Material" {
				p.World.Objects = append(p.World.Objects, Object{p.Inventory[p.Control.Action.SelectedItem].Name, 10, Position{p.Position.X + 5, p.Position.Y + 30}, p.Inventory[p.Control.Action.SelectedItem].Width, p.Inventory[p.Control.Action.SelectedItem].Height})
				p.Inventory = append(p.Inventory[:p.Control.Action.SelectedItem], p.Inventory[p.Control.Action.SelectedItem+1:]...)
			}
		}
		if p.Face == "Right" {
			if p.Inventory[p.Control.Action.SelectedItem].Type == "Material" {
				p.World.Objects = append(p.World.Objects, Object{p.Inventory[p.Control.Action.SelectedItem].Name, 10, Position{p.Position.X + 25, p.Position.Y + 5}, p.Inventory[p.Control.Action.SelectedItem].Width, p.Inventory[p.Control.Action.SelectedItem].Height})
				p.Inventory = append(p.Inventory[:p.Control.Action.SelectedItem], p.Inventory[p.Control.Action.SelectedItem+1:]...)
			}
		}
		if p.Face == "Left" {
			if p.Inventory[p.Control.Action.SelectedItem].Type == "Material" {

				p.World.Objects = append(p.World.Objects, Object{p.Inventory[p.Control.Action.SelectedItem].Name, 10, Position{p.Position.X - 15, p.Position.Y + 5}, p.Inventory[p.Control.Action.SelectedItem].Width, p.Inventory[p.Control.Action.SelectedItem].Height})
				p.Inventory = append(p.Inventory[:p.Control.Action.SelectedItem], p.Inventory[p.Control.Action.SelectedItem+1:]...)
			}
		}
	}
}
