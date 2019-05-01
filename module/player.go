package module

import (
	"crypto/rand"
	"log"
	"math/big"
)

type Player struct {
	ID              string
	Position        Position
	Velocity        Velocity
	Control         Control
	Class           string
	Attributes      Attributes
	BonusAttributes BonusAttributes
	Inventory       []Item
}

type Position struct {
	X float64
	Y float64
}

type Velocity struct {
	X float64
	Y float64
}

type Control struct {
	Right  bool
	Left   bool
	Up     bool
	Down   bool
	Attack Attack
}

type Attributes struct {
	Attack    int
	Intellect int
	Defense   int
}

type BonusAttributes struct {
	Attack    int
	Intellect int
	Defense   int
}

type Item struct {
	Name      string
	Attack    int
	Intellect int
	Defense   int
}
type Attack struct {
	Basic   bool
	Range   bool
	Special bool
}

func (p *Player) UpdatePosition() {

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
