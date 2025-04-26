package system

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jtbonhomme/ebitenecs/component"
	"github.com/jtbonhomme/ebitenecs/entity"
)

var (
	id ID
)

// ID is a type that represents a unique identifier for a system.
type ID int

// String returns the string representation of the ID.
func (id ID) String() string {
	return strconv.Itoa(int(id))
}

// AssignID is a function that assigns a unique ID to a system.
func AssignID() ID {
	id++
	return id
}

// System is an interface that represents a system in the ECS architecture.
type System interface {
	ID() ID
}

// Updater is an interface that represents a system that updates entities in the ECS architecture.
type Updater interface {
	System
	Update(entity.ID, []component.Component, map[entity.ID][]component.Component) error
}

// Drawer is an interface that represents a system that draws entities in the ECS architecture.
type Drawer interface {
	System
	Draw(*ebiten.Image, []component.Component)
}
