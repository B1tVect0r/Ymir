package ecs

import "fmt"

// componentEntityMap contains relational information that allows quick lookup of which entities
// hold a given component
type componentEntityMap map[ComponentID]componentEntitySet

func (cem componentEntityMap) add(cID ComponentID, eID EntityHandle) error {
	if _, ok := cem[cID]; !ok {
		cem[cID] = componentEntitySet{}
	} else if _, ok = cem[cID][eID]; ok {
		return fmt.Errorf("entity %s is already associated with component %s", eID, cID)
	}

	cem[cID][eID] = struct{}{}
	return nil
}

func (cem componentEntityMap) remove(cID ComponentID, eID EntityHandle) {
	delete(cem[cID], eID)
}
