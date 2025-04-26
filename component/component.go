// Package component provides a way to create and manage components in an ECS (Entity-Component-System) architecture.
package component

import (
	"fmt"
	"reflect"
)

// Component is an interface that represents a component in the ECS architecture.
type Component interface {
	Data() interface{}
}

type component struct {
	data interface{}
}

// New creates a new component with the given data.
func New(data interface{}) Component {
	// todo: add a check to control data is a pointer to a struct
	return &component{
		data: data,
	}
}

// Data returns the data of the component.
func (c *component) Data() interface{} {
	return c.data
}

// QueryComponents is a function that queries components matching the given component types from the ECS architecture.
func QueryComponents(c []Component, components ...interface{}) {
	for _, component := range components {
		componentValue := reflect.ValueOf(component)
		// the component to be assigned needs to be a reference to a concrete object
		if componentValue.Kind() != reflect.Ptr {
			panic(fmt.Sprintf("received entity component %s must be a pointer", componentValue.Type().Name()))
		}

		componentValueElem := componentValue.Elem()
		if componentValueElem.Kind() != reflect.Ptr {
			panic(fmt.Sprintf("received entity component %s must be a pointer to pointer", componentValue.Type().Name()))
		}

		componentValueType := componentValueElem.Type()

		for _, rComponent := range c {
			rComponentValue := reflect.ValueOf(rComponent.Data())

			if rComponentValue.Kind() != reflect.Ptr {
				panic(fmt.Sprintf("registered entity component %s must be a pointer", rComponentValue.Type().Name()))
			}

			rComponentValueType := rComponentValue.Type()

			if rComponentValueType == componentValueType {
				componentValueElem.Set(rComponentValue)
				break
			}
		}
	}
}
