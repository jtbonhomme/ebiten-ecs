/*
Package ebitenecs implements the Entity-Component-System (ECS) architecture
for use with the Ebiten game library (https://ebiten.org/). It allows developers to create and manage entities, components, and systems in a modular and flexible way, enabling complex behaviors through simple components and systems.

It provides a framework for managing entities, components, and systems in a game or simulation.
The ECS architecture allows for a flexible and modular design, enabling developers to create complex behaviors by composing simple components and systems.
This package provides the core functionality for creating and managing entities, components, and systems.
It includes methods for registering and unregistering entities and components, as well as for updating and drawing the entities and their components.

# ECS Architecture

More information about the ECS architecture can be found here: https://en.wikipedia.org/wiki/Entity_component_system.

Entites are the basic building blocks of the ECS architecture. They represent individual objects in the game world, such as players, enemies, or items. Each entity can have multiple components associated with it, which define its properties and behaviors.

Components are the data containers in the ECS architecture. They hold the data and state of an entity, such as its position, velocity, or health. Components are typically simple structs that contain only data, without any behavior or logic.

Systems are the logic and behavior of the ECS architecture. They operate on entities and their components, updating their state and performing actions based on the data in the components. Systems are typically implemented as functions or methods that take entities and their components as arguments and perform operations on them.

# Ebitenecs ECS Implementation

In ebitenecs, systems are represented by the Updater and Drawer interfaces, which define the methods for updating and drawing entities and their components.

The Updater interface defines the Update method, which is called every frame to update the state of the entities and their components.

The Drawer interface defines the Draw method, which is called every frame to draw the entities and their components on the screen.

# Usage

First, create an instance of the ECS:

	world := ebitenecs.New()

Then create an entity and register it with the ECS. Don't forget to add a component to the entity:

	countDown := entity.New()
	world.RegisterEntity(
		countDown,
		component.New(
			&CounterComponent{
				Value: 1000000,
			},
		),
	)

Then create a system and register it with the ECS:

	// create a system to manage the CounterComponent
	counterSystem := &CounterSystem{
		id: system.AssignID(),
	}

	// register it in the ECS world as an updater associated with the entity countDown
	world.RegisterUpdater(
		counterSystem,
		countDown,
	)

	// register it in the ECS world as a drawerr associated with the entity countDown
	world.RegisterDrawer(
		counterSystem,
		254,
		countDown,
	)
*/
package ebitenecs
