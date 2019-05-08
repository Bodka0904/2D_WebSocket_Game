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
	BonusAttributes BonusAttributes
	Inventory       []Item
	World           *World
}

func (p *Player) UpdatePlayer() {

	if p.Control.Up {
		p.Position.Y -= p.Velocity.Y
	}
	if p.Control.Down {
		p.Position.Y += p.Velocity.Y
	}
	if p.Control.Left {
		p.Position.X -= p.Velocity.X
	}
	if p.Control.Right {
		p.Position.X += p.Velocity.X
	}
	p.ChangeWorld(WorldList[1])

	if p.Control.Action.Attack {
		if p.Energy >= 0.5 {
			p.Energy -= 0.5
		} else {
			p.Control.Action.Attack = false
		}
	}
	if p.Control.Action.Mine {
		if p.Energy >= 1 {
			p.Energy -= 1
		} else {
			p.Control.Action.Mine = false
		}
	}
	if p.Control.Action.Build {
		if p.Energy >= 3 {
			p.Energy -= 3
		} else {
			p.Control.Action.Build = false
		}
	}
	if p.Energy < 100 {
		p.Energy++
	}

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
	if p.Control.Action.Pick{		
		if p.Position.X <= 
	}
}
