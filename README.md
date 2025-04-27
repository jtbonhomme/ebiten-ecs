# ebiten-ecs

[![Go Reference](https://pkg.go.dev/badge/github.com/jtbonhomme/ebiten-ecs)](https://pkg.go.dev/github.com/jtbonhomme/ebiten-ecs)
[![issues](https://img.shields.io/github/issues/jtbonhomme/ebiten-ecs)](https://github.com/jtbonhomme/ebiten-ecs/issues)
![GitHub Release](https://img.shields.io/github/v/release/jtbonhomme/ebiten-ecs)
[![license](https://img.shields.io/github/license/jtbonhomme/ebiten-ecs)](https://github.com/jtbonhomme/ebiten-ecs/blob/main/LICENSE)
`ebiten-ecs` is a Go library that provides an Entity-Component-System (ECS) framework tailored for use with the [Ebiten](https://ebiten.org/) game library. It simplifies the development of complex 2D games by organizing game logic into entities, components, and systems.

## Features

- **Entity-Component-System Architecture**: Decouple game logic for better maintainability and scalability.
- **Integration with Ebiten**: Seamlessly integrates with Ebiten's rendering and update loops.
- **Lightweight and Performant**: Designed to be efficient and easy to use for small to medium-sized games.
- **Flexible Component Management**: Add, remove, and query components dynamically.

## Installation

To install the library, use:

```bash
go get github.com/jtbonhomme/ebiten-ecs
```


## Running the Example Program

An example program demonstrating the usage of `ebiten-ecs` is provided in the `example` directory. To run the example:

1. Navigate to the `example` directory:
   ```bash
   cd example
   ```

2. Run the example program:
   ```bash
   go run main.go
   ```

This will launch a window showcasing a simple ECS usage (counter).

## Usage

Below is a simple example of how to use `ebiten-ecs` to create a game with entities and components:

```go
package main

import (
	"github.com/jtbonhomme/ebiten-ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world *ecs.World
}

func (g *Game) Update() error {
	// Update all systems in the ECS world
	g.world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render all systems in the ECS world
	g.world.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	// Create a new ECS world
	world := ecs.NewWorld()

	// Add entities, components, and systems to the world
	// ...existing code...

	// Initialize the game
	game := &Game{world: world}

	// Run the game
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebiten ECS Example")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
```

## Why Use `ebiten-ecs`?

- **Simplified Game Logic**: Focus on individual components and systems without worrying about the overall structure.
- **Reusability**: Components and systems can be reused across different projects.
- **Scalability**: Easily extend your game by adding new entities, components, or systems.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the library.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

