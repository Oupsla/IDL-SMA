package sma

import (
	"fmt"
	"os"
	"time"

	"github.com/Oupsla/IDL-SMA/environment"
	ui "github.com/gizak/termui"
	"github.com/spf13/viper"
)

var env *environment.Environment
var showGrid bool
var canvasX int
var canvasY int
var nbTicks int
var delay int
var trace bool

// initConfig : Initalize the system
func initConfig() {

	// Load config with viper
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found...")
		os.Exit(1)
	}

	fmt.Println("\nConfig found !")
	gridSizeX := viper.GetInt("grid.x")
	gridSizeY := viper.GetInt("grid.Y")
	gridTorrique := viper.GetBool("grid.torrique")
	nbParticles := viper.GetInt("nbParticles")
	seed := viper.GetInt64("seed")
	canvasX = viper.GetInt("canvasSize.x")
	canvasY = viper.GetInt("canvasSize.Y")
	showGrid = viper.GetBool("showGrid")
	delay = viper.GetInt("delay")
	nbTicks = viper.GetInt("nbTicks")
	trace = viper.GetBool("trace")

	if trace {
		fmt.Println("gridSizeX = ", gridSizeX)
		fmt.Println("gridSizeY = ", gridSizeY)
		fmt.Println("gridTorrique = ", gridTorrique)
		fmt.Println("canvasX = ", canvasX)
		fmt.Println("canvasY = ", canvasY)
		fmt.Println("delay (in ms) = ", delay)
		fmt.Println("nbTicks (0 = infinite) = ", nbTicks)
		fmt.Println("nbParticles = ", nbParticles)
		fmt.Println(" ")
	}

	// Verify grid size
	if gridSizeX*gridSizeY < nbParticles {
		fmt.Println("Grid is too small for the number of particles ")
		os.Exit(1)
	}

	env, _ = environment.CreateEnvironment(gridSizeX, gridSizeY, nbParticles, seed, gridTorrique, trace)
}

// initGUI: initiatize the console gui
func initGUI() {

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	p := ui.NewPar("Press q to exit")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Particule System"
	p.BorderFg = ui.ColorCyan

	rows := [][]string{
		[]string{"X", " ", " "},
		[]string{" ", "X", " "},
		[]string{" ", " ", "X"},
	}

	table := ui.NewTable()
	table.Rows = rows
	table.FgColor = ui.ColorWhite
	table.BgColor = ui.ColorDefault
	table.TextAlign = ui.AlignCenter

	table.Analysis()
	table.SetSize()
	table.Y = 3
	table.X = 0
	table.Border = false
	table.Seperator = false

	ui.Render(table, p) // feel free to call Render, it's async and non-block
	// handle key q pressing
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/C-x", func(ui.Event) {
		// press ctrl c to quit
		ui.StopLoop()
	})
	ui.Loop() // block until StopLoop is called

}

// Run : run the world
func Run() {
	fmt.Println("... Init Config ...")
	initConfig()
	fmt.Println("... Init GUI ...")
	//initGUI()

	env.Show()
	time.Sleep(time.Duration(delay) * time.Millisecond)

	for i := 0; i < nbTicks; i++ {
		env.Decide()
		env.Show()
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

}
