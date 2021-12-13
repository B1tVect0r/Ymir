package ecs

import (
	"fmt"
	"sync/atomic"
)

type EntityID uint32

const InvalidEntityID EntityID = 0

var lastEntityID = InvalidEntityID

func nextEntityID() EntityID {
	return EntityID(atomic.AddUint32((*uint32)(&lastEntityID), 1))
}

func NewEntity(r *Registry, component Component, components ...Component) (EntityID, error) {

	return r.RegisterEntity(append([]Component{component}, components...)...)
}

func AddComponents(r *Registry, eID EntityID, components ...Component) error {
	return r.AddComponentsToEntity(eID, components...)
}

func RemoveComponents(r *Registry, eID EntityID, cIDs ...ComponentID) error {
	if empty, err := r.RemoveComponentsFromEntity(eID, cIDs...); err != nil {
		return fmt.Errorf("failed to remove components from entity: %w", err)
	} else if empty {
		if err = r.UnregisterEntity(eID); err != nil {
			return fmt.Errorf("failed to unregister entity: %w", err)
		}
	}

	return nil
}

func Destroy(r *Registry, eID EntityID) error {
	return r.UnregisterEntity(eID)
}
