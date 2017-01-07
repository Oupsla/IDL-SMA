package agent

import "fmt"

// Agent : represent a particle with a color and a position
type Agent struct {
	Identifier string
	Color      Color
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

// CreateAgent : Create a new agent with specific color
func CreateAgent(identifier string, color Color, x int, y int) (*Agent, error) {
	return &Agent{
		Identifier: identifier,
		Color:      color,
		X:          x,
		Y:          y,
	}, nil
}

func (agent *Agent) String() string {
	return fmt.Sprintf("Agent with color : %s", agent.Color)
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
	"Blue",
	"Red",
	"Yellow",
	"Green",
	"White",
}

// String returns the English name of the color
func (m Color) String() string { return colors[m-1] }
