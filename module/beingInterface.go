package module

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
	Action Action
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
	Level     int
	Position  Position
}
type Action struct {
	Attack bool
	Mine   bool
	Build  bool
}
