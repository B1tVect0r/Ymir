package ecs

import "fmt"

type entityMapEntry struct {
	e          *Entity
	components entityComponentMap
}

type entityMap map[EntityHandle]*entityMapEntry

func (em entityMap) add(e *Entity, components ...Component) (EntityHandle, error) {
	eID := EntityHandle(e.GetID())

	if _, ok := em[eID]; ok {
		return "", fmt.Errorf("map already contains entity %s", eID)
	}

	entry := &entityMapEntry{
		e:          e,
		components: make(entityComponentMap, len(components)),
	}

	if err := entry.components.add(components...); err != nil {
		return "", fmt.Errorf("failed to add components: %w", err)
	}

	em[eID] = entry
	return eID, nil
}

func (em entityMap) remove(eID EntityHandle) []ComponentID {
	entry, ok := em[eID]
	if !ok {
		return nil
	}

	cIDs := make([]ComponentID, 0, len(entry.components))
	for cID := range entry.components {
		cIDs = append(cIDs, cID)
	}

	delete(em, eID)
	return cIDs
}

func (em entityMap) addComponents(eID EntityHandle, components ...Component) error {
	entry, ok := em[eID]
	if !ok {
		return fmt.Errorf("entity %s does not exist in map", eID)
	}

	if err := entry.components.add(components...); err != nil {
		return fmt.Errorf("failed to add components to entity %s: %w", eID, err)
	}

	return nil
}

func (em entityMap) removeComponents(eID EntityHandle, cIDs ...ComponentID) {
	entry, ok := em[eID]
	if !ok {
		return
	}

	entry.components.remove(cIDs...)
}

func (em entityMap) entityIsEmpty(eID EntityHandle) (bool, error) {
	entry, ok := em[eID]
	if !ok {
		return false, fmt.Errorf("entity %s does not exist in map", eID)
	}

	return len(entry.components) == 0, nil
}

func (em entityMap) numComponents(eID EntityHandle) int {
	entry, ok := em[eID]
	if !ok {
		return 0
	}

	return len(entry.components)
}

func (em entityMap) getComponents(eID EntityHandle, cIDs ...ComponentID) ([]Component, error) {
	entry, ok := em[eID]
	if !ok {
		return nil, fmt.Errorf("entity %s does not exist in map", eID)
	}

	outComponents := make([]Component, 0, len(cIDs))

	for _, cID := range cIDs {
		c, ok := entry.components[cID]
		if !ok {
			return nil, fmt.Errorf("failed to find component %s on entity %s", cID, eID)
		}
		outComponents = append(outComponents, c)
	}

	return outComponents, nil
}
