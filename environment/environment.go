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
			randomDirection := randomGen.Intn(8)

			if grid[y1][x1] == nil {

				a, _ := agent.CreateAgent(strconv.Itoa(i), agent.RandomColor(), x1, y1, randomDirection)

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

func (env *Environment) movingAgent(currentAgent *agent.Agent) {

	// Memorize old position
	oldX := currentAgent.X
	oldY := currentAgent.Y

	// Remove in the grid
	env.Grid[currentAgent.Y][currentAgent.X] = nil

	// Set new position
	env.newPosition(currentAgent)

	// Testing borders
	if currentAgent.X >= env.GridX {
		if env.Trace {
			fmt.Printf("%s collission with East \n", currentAgent.Identifier)
		}
		if env.Torrique {
			currentAgent.X = 0
		} else {
			currentAgent.X = env.GridX - 2
		}

	}

	if currentAgent.X < 0 {
		if env.Trace {
			fmt.Printf("%s collission with West \n", currentAgent.Identifier)
		}
		if env.Torrique {
			currentAgent.X = env.GridX - 1
		} else {
			currentAgent.X = 1
		}
	}

	if currentAgent.Y >= env.GridY {
		if env.Trace {
			fmt.Printf("%s collission with South \n", currentAgent.Identifier)
		}
		if env.Torrique {
			currentAgent.Y = 0
		} else {
			currentAgent.Y = env.GridY - 1
		}
	}

	if currentAgent.Y < 0 {
		if env.Trace {
			fmt.Printf("%s collission with North \n", currentAgent.Identifier)
		}
		if env.Torrique {
			currentAgent.Y = env.GridY - 1
		} else {
			currentAgent.Y = 1
		}
	}

	// Testing collisionss
	if env.Grid[currentAgent.Y][currentAgent.X] != nil {
		if env.Trace {
			fmt.Printf("%s collission with %s \n", currentAgent.Identifier, env.Grid[currentAgent.Y][currentAgent.X].Identifier)
		}

		// Changing Direction
		oldDirection := currentAgent.Direction
		currentAgent.Direction = env.Grid[currentAgent.Y][currentAgent.X].Direction
		env.Grid[currentAgent.Y][currentAgent.X].Direction = oldDirection

		// Resetting position
		env.Grid[oldY][oldX] = currentAgent
		currentAgent.X = oldX
		currentAgent.Y = oldY

		// Moving agent again
		env.movingAgent(currentAgent)

	} else {
		env.Grid[currentAgent.Y][currentAgent.X] = currentAgent
	}

}

func (env *Environment) newPosition(currentAgent *agent.Agent) {

	switch currentAgent.Direction {
	// North
	case 0:
		currentAgent.Y--
		if env.Trace {
			fmt.Printf("%s moving North \n", currentAgent.Identifier)
		}
	// North East
	case 1:
		currentAgent.X++
		currentAgent.Y--
		if env.Trace {
			fmt.Printf("%s moving North East \n", currentAgent.Identifier)
		}
	// East
	case 2:
		currentAgent.X++
		if env.Trace {
			fmt.Printf("%s moving East \n", currentAgent.Identifier)
		}
	// South East
	case 3:
		currentAgent.Y++
		currentAgent.X++
		if env.Trace {
			fmt.Printf("%s moving South East \n", currentAgent.Identifier)
		}
	// South
	case 4:
		currentAgent.Y++
		if env.Trace {
			fmt.Printf("%s moving South \n", currentAgent.Identifier)
		}
	// South West
	case 5:
		currentAgent.Y++
		currentAgent.X--
		if env.Trace {
			fmt.Printf("%s moving South West \n", currentAgent.Identifier)
		}
	// West
	case 6:
		currentAgent.X--
		if env.Trace {
			fmt.Printf("%s moving West \n", currentAgent.Identifier)
		}
	// North West
	case 7:
		currentAgent.Y--
		currentAgent.X--
		if env.Trace {
			fmt.Printf("%s moving North West \n", currentAgent.Identifier)
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

				lineResult += env.Grid[i][j].String()
			}
		}
		fmt.Printf(lineResult + "\n")
	}
}
