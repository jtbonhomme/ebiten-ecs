package ecs

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/jtbonhomme/ebiten-ecs/component"
	"github.com/jtbonhomme/ebiten-ecs/entity"
	"github.com/jtbonhomme/ebiten-ecs/system"
)

const (
	MaxEntities int = 512
	MaxSystems  int = 256
	MaxDrawers  int = 256
)

// ECS is the main structure for the Entity-Component-System architecture.
// It provides methods to register and unregister entities, components, updaters, and drawers.
type ECS struct {
	updaters           []system.Updater
	drawers            map[int][]system.Drawer
	entitiesRegistry   map[system.ID][]entity.Entity
	componentsRegistry map[entity.ID][]component.Component
}

// New creates a new ECS instance with initialized registries for entities and components.
// It also initializes the updaters and drawers slices.
// The ECS instance is ready to be used for managing entities and their components.
// It is important to note that the ECS instance should be used in a single-threaded context
// to avoid concurrent access issues.
// The ECS instance is not thread-safe, and concurrent access to its methods may lead to undefined behavior.
// It is recommended to use a single goroutine to manage the ECS instance and its entities.
// This ensures that the ECS instance is used in a safe and predictable manner.
func New() *ECS {
	return &ECS{
		updaters:           []system.Updater{},
		drawers:            make(map[int][]system.Drawer, MaxDrawers),
		entitiesRegistry:   make(map[system.ID][]entity.Entity, MaxSystems),
		componentsRegistry: make(map[entity.ID][]component.Component, MaxEntities),
	}
}

// RegisterEntity registers an entity and its components in the ECS.
// It takes an entity and a variadic number of components as arguments.
// The entity is assigned a unique ID, and the components are associated with the entity.
// The components are stored in the components registry, which maps entity IDs to their respective components.
// The method checks if the components are pointers to structs, and panics if they are not.
func (ecs *ECS) RegisterEntity(e entity.Entity, components ...component.Component) {
	for _, component := range components {
		// check component data member is a ptr
		componentValue := reflect.ValueOf(component.Data())

		if componentValue.Kind() != reflect.Ptr {
			panic(fmt.Sprintf("the entity component %q you are trying to register MUST be a pointer", componentValue.Type().Name()))
		}

		ecs.componentsRegistry[e.ID()] = append(ecs.componentsRegistry[e.ID()], component)
	}
}

func deleteFromSlice(l []entity.Entity, i int) []entity.Entity {
	if len(l) == 0 || i >= len(l) || i < 0 {
		return l
	}
	l[i], l[len(l)-1] = l[len(l)-1], l[i]
	return l[:len(l)-1]
}

// UnregisterEntity removes an entity and its components from the ECS.
// It takes an entity ID as an argument and removes the entity from the entities registry.
// It also removes the components associated with the entity from the components registry.
// The method iterates through the entities registry and removes the entity from the list of entities
// associated with the system ID. It also deletes the components associated with the entity ID from the components registry.
func (ecs *ECS) UnregisterEntity(id entity.ID) {
	for sid, entities := range ecs.entitiesRegistry {
		for i, e := range entities {
			if e.ID() == id {
				ecs.entitiesRegistry[sid] = deleteFromSlice(entities, i)
			}
		}
	}

	delete(ecs.componentsRegistry, id)
}

// UnregisterSystem removes a system and its associated entities from the ECS.
// It takes a system ID as an argument and removes the system from the entities registry.
func (ecs *ECS) RegisterUpdater(s system.Updater, e ...entity.Entity) {
	ecs.updaters = append(ecs.updaters, s)
	ecs.entitiesRegistry[s.ID()] = append(ecs.entitiesRegistry[s.ID()], e...)
}

// UnregisterSystem removes a system and its associated entities from the ECS.
// It takes a system ID as an argument and removes the system from the entities registry.
func (ecs *ECS) RegisterDrawer(s system.Drawer, zIndex int, e ...entity.Entity) {
	_, ok := ecs.drawers[zIndex]
	if !ok {
		ecs.drawers[zIndex] = []system.Drawer{}
	}

	ecs.drawers[zIndex] = append(ecs.drawers[zIndex], s)
	ecs.entitiesRegistry[s.ID()] = append(ecs.entitiesRegistry[s.ID()], e...)
}

// UnregisterSystem removes a system and its associated entities from the ECS.
// It takes a system ID as an argument and removes the system from the entities registry.
func (ecs *ECS) QueryEntityComponents(e entity.Entity, components ...interface{}) {
	registeredComponents := ecs.componentsRegistry[e.ID()]
	component.QueryComponents(registeredComponents, components...)
}

// FilterEntities filters the entities associated with a system.
// It takes a system as an argument and returns a slice of entities associated with the system ID.
func (ecs *ECS) FilterEntities(s system.System) []entity.Entity {
	return ecs.entitiesRegistry[s.ID()]
}

// Updaters returns the slice of registered updaters in the ECS.
func (ecs *ECS) Updaters() []system.Updater {
	return ecs.updaters
}

// Drawers returns the map of registered drawers in the ECS.
func (ecs *ECS) Drawers() map[int][]system.Drawer {
	return ecs.drawers
}

// Update iterates through the registered updaters and updates the entities associated with them.
func (ecs *ECS) Update() error {
	for _, s := range ecs.Updaters() {
		for _, e := range ecs.FilterEntities(s) {
			registeredComponents := ecs.componentsRegistry[e.ID()]
			err := s.Update(e.ID(), registeredComponents, ecs.componentsRegistry)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Draw iterates through the registered drawers and draws the entities associated with them.
func (ecs *ECS) Draw(screen *ebiten.Image) {
	// https://go.dev/blog/maps - Iteration order
	drawers := ecs.Drawers()

	zIndexes := make([]int, 0, len(drawers))

	for i := range drawers {
		zIndexes = append(zIndexes, i)
	}
	sort.Ints(zIndexes)

	for _, i := range zIndexes {
		for _, d := range drawers[i] {
			for _, e := range ecs.FilterEntities(d) {
				registeredComponents := ecs.componentsRegistry[e.ID()]
				d.Draw(screen, registeredComponents)
			}
		}
	}
}
