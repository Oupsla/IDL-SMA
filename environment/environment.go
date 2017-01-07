package environment

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Oupsla/IDL-SMA/agent"
)

// Environment : represent the grid with agents
type Environment struct {
	GridX           int
	GridY           int
	Torrique        bool
	Trace           bool
	Grid            [][]*agent.Agent
	AgentList       []*agent.Agent
	SourceRandom    rand.Source
	GeneratorRandom *rand.Rand
}

// NewEnvironement : Empty constructor
func NewEnvironement() (*Environment, error) {
	return &Environment{}, nil
}

// CreateEnvironment : Create a new environment and place agent
func CreateEnvironment(gridX int, gridY int, nbAgent int, randomSeed int64, torrique bool, trace bool) (*Environment, error) {

	// Create random generator with the seed
	sourceRandom := rand.NewSource(randomSeed)
	randomGen := rand.New(sourceRandom)

	// Allocate the grid
	grid := make([][]*agent.Agent, gridY) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range grid {
		grid[i] = make([]*agent.Agent, gridX)
	}

	// Allocate the list
	agentList := make([]*agent.Agent, nbAgent)

	// Create Agents
	for i := 0; i < nbAgent; i++ {

		// Infinite loop till we place the agent
		for {
			x1 := randomGen.Intn(gridX)
			y1 := randomGen.Intn(gridY)

			if grid[y1][x1] == nil {

				a, _ := agent.CreateAgent(strconv.Itoa(i), agent.Blue, x1, y1)

				grid[y1][x1] = a
				agentList[i] = a

				break
			}
		}

	}

	return &Environment{
		Grid:            grid,
		AgentList:       agentList,
		GridX:           gridX,
		GridY:           gridY,
		Torrique:        torrique,
		SourceRandom:    sourceRandom,
		GeneratorRandom: randomGen,
		Trace:           trace,
	}, nil
}

// Decide : Moving agents
func (env *Environment) Decide() {

	//Shuffle list agent
	fmt.Println("\n... Shuffle list of agents ... ")
	for m := range env.AgentList {
		n := rand.Intn(m + 1)
		env.AgentList[m], env.AgentList[n] = env.AgentList[n], env.AgentList[m]
	}

	fmt.Println("... Moving Agents ...\n ")
	for m := range env.AgentList {
		env.movingAgent(env.AgentList[m])
	}

}

func (env *Environment) movingAgent(agent *agent.Agent) {

	// Memorize old position
	oldX := agent.X
	oldY := agent.Y

	// Remove in the grid
	env.Grid[agent.Y][agent.X] = nil

	// Get new position
	randomDirection := env.GeneratorRandom.Intn(8)
	env.newPosition(agent, randomDirection)

	// Testing borders
	if agent.X >= env.GridX {
		if env.Trace {
			fmt.Printf("%s collission with East \n", agent.Identifier)
		}
		if env.Torrique {
			agent.X = 0
		} else {
			agent.X = env.GridX - 2
		}

	}

	if agent.X < 0 {
		if env.Trace {
			fmt.Printf("%s collission with West \n", agent.Identifier)
		}
		if env.Torrique {
			agent.X = env.GridX - 1
		} else {
			agent.X = 1
		}
	}

	if agent.Y >= env.GridY {
		if env.Trace {
			fmt.Printf("%s collission with South \n", agent.Identifier)
		}
		if env.Torrique {
			agent.Y = 0
		} else {
			agent.Y = env.GridY - 1
		}
	}

	if agent.Y < 0 {
		if env.Trace {
			fmt.Printf("%s collission with North \n", agent.Identifier)
		}
		if env.Torrique {
			agent.Y = env.GridY - 1
		} else {
			agent.Y = 1
		}
	}

	// Testing collisionss
	if env.Grid[agent.Y][agent.X] != nil {
		if env.Trace {
			fmt.Printf("%s collission \n", agent.Identifier)
		}
		env.Grid[oldY][oldX] = agent
		agent.X = oldX
		agent.Y = oldY
	} else {
		env.Grid[agent.Y][agent.X] = agent
	}

}

func (env *Environment) newPosition(agent *agent.Agent, direction int) {

	switch direction {
	// North
	case 0:
		agent.Y--
		if env.Trace {
			fmt.Printf("%s moving North \n", agent.Identifier)
		}
	// North East
	case 1:
		agent.X++
		agent.Y--
		if env.Trace {
			fmt.Printf("%s moving North East \n", agent.Identifier)
		}
	// East
	case 2:
		agent.X++
		if env.Trace {
			fmt.Printf("%s moving East \n", agent.Identifier)
		}
	// South East
	case 3:
		agent.Y++
		agent.X++
		if env.Trace {
			fmt.Printf("%s moving South East \n", agent.Identifier)
		}
	// South
	case 4:
		agent.Y++
		if env.Trace {
			fmt.Printf("%s moving South \n", agent.Identifier)
		}
	// South West
	case 5:
		agent.Y++
		agent.X--
		if env.Trace {
			fmt.Printf("%s moving South West \n", agent.Identifier)
		}
	// West
	case 6:
		agent.X--
		if env.Trace {
			fmt.Printf("%s moving West \n", agent.Identifier)
		}
	// North West
	case 7:
		agent.Y--
		agent.X--
		if env.Trace {
			fmt.Printf("%s moving North West \n", agent.Identifier)
		}
	}

}

// Show : Print the actual env
func (env *Environment) Show() {
	/* output each array element's value */
	for i := 0; i < env.GridY; i++ {
		lineResult := ""
		for j := 0; j < env.GridX; j++ {

			if env.Grid[i][j] == nil {
				lineResult += "."
			} else {

				lineResult += env.Grid[i][j].Identifier
			}
		}
		fmt.Printf(lineResult + "\n")
	}
}
