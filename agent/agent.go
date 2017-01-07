package agent

import (
	"fmt"
	"math/rand"
	"time"
)

// Agent : represent a particle with a color and a position
type Agent struct {
	Color        Color
	SourceRandom rand.Source
	RandomGen    *rand.Rand
	X            int
	Y            int
}

// NewAgent : Empty constructor
func NewAgent() (*Agent, error) {

	sourceRandom := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(sourceRandom)

	return &Agent{
		Color:        Blue,
		SourceRandom: sourceRandom,
		RandomGen:    randomGen,
		X:            0,
		Y:            0,
	}, nil
}

// CreateAgent : Create a new agent with specific color and seed
func CreateAgent(color Color, seed int64, x int, y int) (*Agent, error) {

	sourceRandom := rand.NewSource(seed)
	randomGen := rand.New(sourceRandom)

	return &Agent{
		Color:        color,
		SourceRandom: sourceRandom,
		RandomGen:    randomGen,
		X:            x,
		Y:            y,
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

	agent.move(agent.RandomGen.Intn(8))

}

func (agent *Agent) move(direction int) {

	switch direction {
	// North
	case 0:
		agent.Y--
	// North East
	case 1:
		agent.X++
		agent.Y--
	// East
	case 2:
		agent.X++
	// South East
	case 3:
		agent.Y++
		agent.X++
	// South
	case 4:
		agent.Y++
	// South West
	case 5:
		agent.Y++
		agent.X--
	// West
	case 6:
		agent.X--
	// North West
	case 7:
		agent.Y--
		agent.X--
	}

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
