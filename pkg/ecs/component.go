package ecs

type ComponentID uint16

type Component interface {
	ID() ComponentID
}
