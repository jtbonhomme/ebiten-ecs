package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/jtbonhomme/ebitenecs"
	"github.com/jtbonhomme/ebitenecs/component"
	"github.com/jtbonhomme/ebitenecs/entity"
	"github.com/jtbonhomme/ebitenecs/system"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// CounterComponent is a simple struct to demonstrate the use of components.
// It contains a single field, Value, which is an integer.
type CounterComponent struct {
	Value int
}

// CounterSystem is a simple system that updates the CounterComponent.
// It implements the Updater interface from the ebitenecs package, and the Drawer interface.
// The Update method decrements the Value field of the CounterComponent by 1.
// The system is identified by a unique ID, which is assigned using the AssignID function.
// The ID method returns the unique ID of the system.
type CounterSystem struct {
	id system.ID
}

// ID returns the unique ID of the CounterSystem.
func (cs *CounterSystem) ID() system.ID {
	return cs.id
}

// Update is called every frame to update the CounterComponent.
// It decrements the Value field of the CounterComponent by 1.
func (cs *CounterSystem) Update(self entity.ID, c []component.Component, r map[entity.ID][]component.Component) error {
	var counter *CounterComponent

	component.QueryComponents(c, &counter)
	counter.Value--

	return nil
}

// Draw is a simple system that draws the CounterComponent value on screen.
func (cs *CounterSystem) Draw(
	screen *ebiten.Image,
	c []component.Component) {
	var counter *CounterComponent

	component.QueryComponents(c, &counter)

	ebitenutil.DebugPrintAt(screen,
		fmt.Sprintf("Counter value is %d", counter.Value),
		320, 240)
}

// Game is the main structure for the game.
type Game struct {
	world *ebitenecs.ECS
}

// Update is called every frame to update the game state.
func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	// update the ECS world
	// this will call the Update method of all registered updaters
	err := g.world.Update()
	if err != nil {
		return err
	}

	return nil
}

// Draw is called every frame to draw the game state on the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 1})
	msg := fmt.Sprintf(`
TPS:            %0.2f
FPS:            %0.2f
PRESS ESCAPE TO QUIT`,
		ebiten.ActualTPS(), ebiten.ActualFPS())

	ebitenutil.DebugPrint(screen, msg)

	// draw the ECS world
	g.world.Draw(screen)
}

// Layout is called every frame to set the screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// main function initializes the game and starts the ebiten loop.
// It sets the window size, title, and creates a new Game instance.
// It also creates a simple counter entity.
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Glow Demo (Ebitengine))")

	// create a new game with ECS world
	g := &Game{
		world: ebitenecs.New(),
	}

	// create a new entity countDown wth a CounterComponent
	// and register it in the ECS world.
	countDown := entity.New()
	g.world.RegisterEntity(
		countDown,
		component.New(
			&CounterComponent{
				Value: 1000000,
			},
		),
	)

	// create a system to manage the CounterComponent
	counterSystem := &CounterSystem{
		id: system.AssignID(),
	}

	// register it in the ECS world as an updater associated with the entity countDown
	g.world.RegisterUpdater(
		counterSystem,
		countDown,
	)

	// register it in the ECS world as a drawerr associated with the entity countDown
	g.world.RegisterDrawer(
		counterSystem,
		254,
		countDown,
	)

	// run the ebiten game loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
