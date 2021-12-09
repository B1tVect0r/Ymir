package ecs

import "fmt"

type entityComponentMap map[ComponentID]Component

func getDuplicateIDs(components ...Component) []ComponentID {
	encountered := make(map[ComponentID]struct{}, len(components))

	outDuplicates := make([]ComponentID, 0, len(components))

	for _, c := range components {
		cID := c.Type()
		if _, ok := encountered[cID]; ok {
			outDuplicates = append(outDuplicates, cID)
		}
		encountered[cID] = struct{}{}
	}

	return outDuplicates
}

func (ecm entityComponentMap) contains(component Component) bool {
	_, ok := ecm[component.Type()]
	return ok
}

func (ecm entityComponentMap) add(components ...Component) error {
	if dupIDs := getDuplicateIDs(components...); len(dupIDs) > 0 {
		return fmt.Errorf("the following components are provided more than once: %v", dupIDs)
	}

	for _, c := range components {
		if ecm.contains(c) {
			return fmt.Errorf("component %s already exists in the map", c.Type())
		}
	}

	for _, c := range components {
		ecm[c.Type()] = c
	}

	return nil
}

func (ecm entityComponentMap) remove(cIDs ...ComponentID) {
	for _, cID := range cIDs {
		delete(ecm, cID)
	}
}
