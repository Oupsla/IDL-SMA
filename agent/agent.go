package agent

import "fmt"

// Agent : represent a particle with a color and a position
type Agent struct {
	Color Color
}

// NewAgent : Empty constructor
func NewAgent() (*Agent, error) {
	return &Agent{
		Color: Blue,
	}, nil
}

// CreateAgent : Create a new agent with specific color
func CreateAgent(color Color) (*Agent, error) {
	return &Agent{
		Color: color,
	}, nil
}

func (agent *Agent) String() string {
	return fmt.Sprintf("Agent with color : %s", agent.Color)
}

// Update : Update the position of the agent
func (agent *Agent) Update() {

}

// Decide : Decide what to do with a limit or collision
func (agent *Agent) Decide() {

}

func (agent *Agent) move(direction string) {

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
