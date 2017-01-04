package agent

type Agent struct {
  Color string
  Position [2]int
}

func NewAgent() (*Agent, error) {
  return &Agent{}, nil
}

func CreateAgent(color string, position [2]int) (*Agent, error) {
  return &Agent{
    Color: color,
    Position: position,
  }, nil
}


func (agent *Agent) Update() {

}

// Process to decide what to do with a limit or collision
func (agent *Agent) Decide() {

}

func (agent *Agent) move(direction string) {

}
