package ecs

import "fmt"

// componentEntityMap contains relational information that allows quick lookup of which entities
// hold a given component
type componentEntityMap map[ComponentID]componentEntitySet

func (cem componentEntityMap) add(cID ComponentID, eID EntityID) error {
	if _, ok := cem[cID]; !ok {
		cem[cID] = componentEntitySet{}
	} else if _, ok = cem[cID][eID]; ok {
		return fmt.Errorf("entity %d is already associated with component %d", eID, cID)
	}

	cem[cID][eID] = struct{}{}
	return nil
}

func (cem componentEntityMap) remove(cID ComponentID, eID EntityID) {
	delete(cem[cID], eID)
}
