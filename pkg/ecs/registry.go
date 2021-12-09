package ecs

import (
	"fmt"
	"sync"
)

type Registry struct {
	sync.RWMutex
	em                  entityMap
	entitiesByComponent componentEntityMap
	systems             map[SystemID]*systemWrapper
}

func (r *Registry) RegisterEntity(e *Entity, cs ...Component) (EntityHandle, error) {
	r.Lock()
	defer r.Unlock()

	eID, err := r.em.add(e, cs...)
	if err != nil {
		return "", fmt.Errorf("failed to add entity to entityMap: %w", err)
	}

	for _, c := range cs {
		if err := r.entitiesByComponent.add(c.Type(), eID); err != nil {
			return "", fmt.Errorf("failed to register entity-component relationship: %w", err)
		}
	}

	return eID, nil
}

func (r *Registry) UnregisterEntity(eID EntityHandle) error {
	r.Lock()
	defer r.Unlock()

	cIDs := r.em.remove(eID)

	for _, cID := range cIDs {
		r.entitiesByComponent.remove(cID, eID)
	}

	return nil
}

func (r *Registry) AddComponentsToEntity(eID EntityHandle, cs ...Component) error {
	r.Lock()
	defer r.Unlock()

	if err := r.em.addComponents(eID, cs...); err != nil {
		return fmt.Errorf("failed to add components to entity %s: %w", eID, err)
	}

	for _, c := range cs {
		if err := r.entitiesByComponent.add(c.Type(), eID); err != nil {
			return fmt.Errorf("failed to register entity-component relationship: %w", err)
		}
	}

	return nil
}

func (r *Registry) RemoveComponentsFromEntity(eID EntityHandle, cIDs ...ComponentID) (bool, error) {
	r.Lock()
	defer r.Unlock()

	r.em.removeComponents(eID, cIDs...)

	for _, cID := range cIDs {
		r.entitiesByComponent.remove(cID, eID)
	}

	return r.em.entityIsEmpty(eID)
}

func (r *Registry) GetComponentsFromEntity(eID EntityHandle, cIDs ...ComponentID) ([]Component, error) {
	r.RLock()
	defer r.RUnlock()

	return r.em.getComponents(eID, cIDs...)
}

func (r *Registry) QueryEntities(qm SystemQueryMode, incl, excl []ComponentID) componentEntitySet {
	r.RLock()
	defer r.RUnlock()

	inclEntitySets := make([]componentEntitySet, 0, len(incl))
	exclEntitySets := make([]componentEntitySet, 0, len(excl))

	for _, inclType := range incl {
		es, ok := r.entitiesByComponent[inclType]
		if !ok {
			return componentEntitySet{}
		}

		inclEntitySets = append(inclEntitySets, es.clone())
	}

	for _, exclType := range excl {
		es, ok := r.entitiesByComponent[exclType]
		if !ok {
			continue
		}

		exclEntitySets = append(exclEntitySets, es.clone())
	}

	// Include only entities that have ALL included component types
	finalIncludeSet := intersect(inclEntitySets...)

	// In explicit mode, if an entity has any components OTHER than the ones we've requested
	// we exclude it from the result set
	if qm == Exclusive {
		for eID := range finalIncludeSet {
			if r.em.numComponents(eID) > len(incl) {
				delete(finalIncludeSet, eID)
			}
		}
	}

	// Exclude entities that have ANY excluded component types
	finalExcludeSet := union(exclEntitySets...)

	return finalIncludeSet.difference(finalExcludeSet)
}

func (r *Registry) registerSystems(sws ...*systemWrapper) {
	r.Lock()
	defer r.Unlock()

	for _, sw := range sws {
		r.systems[sw.id()] = sw
	}
}

func (r *Registry) unregisterSystems(sIDs ...SystemID) {
	r.Lock()
	defer r.Unlock()

	for _, sID := range sIDs {
		delete(r.systems, sID)
	}
}

type RegistryConfig struct {
	InitialEntityCapacity     int
	InitialComponentsCapacity int
	InitialSystemsCapacity    int
}

func NewRegistry() *Registry {
	return NewRegistryWithConfig(RegistryConfig{})
}

func NewRegistryWithConfig(c RegistryConfig) *Registry {
	return &Registry{
		em:                  make(entityMap, c.InitialEntityCapacity),
		entitiesByComponent: make(componentEntityMap, c.InitialComponentsCapacity),
		systems:             make(map[SystemID]*systemWrapper, c.InitialSystemsCapacity),
	}
}
