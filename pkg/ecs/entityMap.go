package ecs

import "fmt"

type entityMap map[EntityID]entityComponentMap

func (em entityMap) add(components ...Component) (EntityID, error) {
	eID := nextEntityID()

	if _, ok := em[eID]; ok {
		return InvalidEntityID, fmt.Errorf("map already contains entity %d", eID)
	}

	ecm := make(entityComponentMap, len(components))

	if err := ecm.add(components...); err != nil {
		return InvalidEntityID, fmt.Errorf("failed to add components: %w", err)
	}

	em[eID] = ecm
	return eID, nil
}

func (em entityMap) remove(eID EntityID) []ComponentID {
	ecm, ok := em[eID]
	if !ok {
		return nil
	}

	cIDs := make([]ComponentID, 0, len(ecm))
	for cID := range ecm {
		cIDs = append(cIDs, cID)
	}

	delete(em, eID)
	return cIDs
}

func (em entityMap) addComponents(eID EntityID, components ...Component) error {
	ecm, ok := em[eID]
	if !ok {
		return fmt.Errorf("entity %d does not exist in map", eID)
	}

	if err := ecm.add(components...); err != nil {
		return fmt.Errorf("failed to add components to entity %d: %w", eID, err)
	}

	return nil
}

func (em entityMap) removeComponents(eID EntityID, cIDs ...ComponentID) {
	ecm, ok := em[eID]
	if !ok {
		return
	}

	ecm.remove(cIDs...)
}

func (em entityMap) entityIsEmpty(eID EntityID) (bool, error) {
	ecm, ok := em[eID]
	if !ok {
		return false, fmt.Errorf("entity %d does not exist in map", eID)
	}

	return len(ecm) == 0, nil
}

func (em entityMap) numComponents(eID EntityID) int {
	ecm, ok := em[eID]
	if !ok {
		return 0
	}

	return len(ecm)
}

func (em entityMap) getComponents(eID EntityID, cIDs ...ComponentID) (entityComponentMap, error) {
	ecm, ok := em[eID]
	if !ok {
		return nil, fmt.Errorf("entity %d does not exist in map", eID)
	}

	outECM := make(entityComponentMap, len(cIDs))

	for _, cID := range cIDs {
		c, ok := ecm[cID]
		if !ok {
			return nil, fmt.Errorf("failed to find component %d on entity %d", cID, eID)
		}
		outECM[cID] = c
	}

	return outECM, nil
}
