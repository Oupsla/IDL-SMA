package environment

import (
	"fmt"
	"math/rand"

	"github.com/Oupsla/IDL-SMA/agent"
)

// Environment : represent the grid with agents
type Environment struct {
	GridX     int
	GridY     int
	Torrique  bool
	Grid      [][]*agent.Agent
	AgentList []*agent.Agent
}

// NewEnvironement : Empty constructor
func NewEnvironement() (*Environment, error) {
	return &Environment{}, nil
}

// CreateEnvironment : Create a new environment and place agent
func CreateEnvironment(gridX int, gridY int, nbAgent int, randomSeed int64, torrique bool) (*Environment, error) {

	// Create random generator with the seed
	sourceRandom := rand.NewSource(randomSeed)
	randomGen := rand.New(sourceRandom)

	// Allocate the grid
	grid := make([][]*agent.Agent, gridX) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range grid {
		grid[i] = make([]*agent.Agent, gridY)
	}

	// Allocate the list
	agentList := make([]*agent.Agent, nbAgent) // One row per unit of y.

	// Create Agents
	for i := 0; i < nbAgent; i++ {
		a, _ := agent.NewAgent()
		agentList[i] = a

		// Infinite loop till we place the agent
		for {
			x1 := randomGen.Intn(gridX)
			y1 := randomGen.Intn(gridY)

			if grid[x1][y1] == nil {
				grid[x1][y1] = a
				break
			}
		}

	}

	return &Environment{
		Grid:      grid,
		AgentList: agentList,
		GridX:     gridX,
		GridY:     gridY,
		Torrique:  torrique,
	}, nil
}

// Show : Print the actual env
func (environment *Environment) Show() {
	/* output each array element's value */
	for i := 0; i < environment.GridX; i++ {
		lineResult := ""
		for j := 0; j < environment.GridY; j++ {
			if environment.Grid[i][j] == nil {
				lineResult += "."
			} else {
				lineResult += "X"
			}
		}
		fmt.Printf(lineResult + "\n")
	}
}

// TO shuffle array : http://stackoverflow.com/questions/12264789/shuffle-array-in-go
