package ecs

import (
	"fmt"

	"github.com/google/uuid"
)

type EntityHandle string

func NewEntity(r *Registry, component Component, components ...Component) (EntityHandle, error) {
	e := &Entity{ID: uuid.NewString()}

	return r.RegisterEntity(e, append([]Component{component}, components...)...)
}

func AddComponents(r *Registry, eID EntityHandle, components ...Component) error {
	return r.AddComponentsToEntity(eID, components...)
}

func RemoveComponents(r *Registry, eID EntityHandle, cIDs ...ComponentID) error {
	if empty, err := r.RemoveComponentsFromEntity(eID, cIDs...); err != nil {
		return fmt.Errorf("failed to remove components from entity: %w", err)
	} else if empty {
		if err = r.UnregisterEntity(eID); err != nil {
			return fmt.Errorf("failed to unregister entity: %w", err)
		}
	}

	return nil
}

func Destroy(r *Registry, eID EntityHandle) error {
	return r.UnregisterEntity(eID)
}
