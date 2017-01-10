package agent

import (
	"fmt"
	"time"
	"math/rand"
)

// Agent : represent a particle with a color and a position
type Agent struct {
	Identifier string
	Color      Color
	Direction  int
	X          int
	Y          int
}

// NewAgent : Empty constructor
func NewAgent() (*Agent, error) {
	return &Agent{
		Identifier: "N",
		Color:      Blue,
		X:          0,
		Y:          0,
	}, nil
}

func RandomColor() (color Color) {
	// Create random generator with the seed
	sourceRandom := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(sourceRandom)
	number := randomGen.Intn(5)

	switch number {
	case 0:
		return Blue
	case 1:
		return Red
	case 2:
		return Yellow
	case 3:
		return Green
	case 4:
		return White
	}

	return Blue
}

// CreateAgent : Create a new agent with specific color
func CreateAgent(identifier string, color Color, x int, y int, direction int) (*Agent, error) {
	return &Agent{
		Identifier: identifier,
		Color:      color,
		X:          x,
		Y:          y,
		Direction:  direction,
	}, nil
}

func (agent *Agent) String() string {
	return fmt.Sprintf("%s", agent.Color)
}

/* ############# Color enum ##############"" */

// Color type
type Color int

const (
	// Blue Color
	Blue Color = 1 + iota
	// Red Color
	Red
	// Yellow Color
	Yellow
	// Green Color
	Green
	// White Color
	White
)

var colors = [...]string{
	"B",
	"R",
	"Y",
	"G",
	"W",
}

// String returns the English name of the color
func (m Color) String() string { return colors[m-1] }
