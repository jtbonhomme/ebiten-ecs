package entity

import "strconv"

var (
	id ID
)

// ID is a type that represents a unique identifier for an entity.
type ID int

// String returns the string representation of the ID.
func (id ID) String() string {
	return strconv.Itoa(int(id))
}

// AssignID is a function that assigns a unique ID to an entity.
func AssignID() ID {
	id++
	return id
}

// Entity is an interface that represents an entity in the ECS architecture.
type Entity interface {
	ID() ID
}

type entity struct {
	id ID
}

// New creates a new entity with a unique ID.
func New() Entity {
	return &entity{
		id: AssignID(),
	}
}

// ID returns the unique ID of the entity.
func (e *entity) ID() ID {
	return e.id
}
