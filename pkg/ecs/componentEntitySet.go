package ecs

import "sort"

// componentEntitySet
type componentEntitySet map[EntityID]struct{}

func (ces componentEntitySet) add(e EntityID) {
	ces[e] = struct{}{}
}

func (ces componentEntitySet) remove(e EntityID) {
	delete(ces, e)
}

func (ces componentEntitySet) contains(e EntityID) bool {
	_, ok := ces[e]
	return ok
}

func (ces componentEntitySet) clone() componentEntitySet {
	outSet := make(componentEntitySet, len(ces))
	for e := range ces {
		outSet.add(e)
	}

	return outSet
}

func (ces componentEntitySet) union(others ...componentEntitySet) componentEntitySet {
	outSet := ces.clone()

	for _, os := range others {
		for e := range os {
			outSet.add(e)
		}
	}

	return outSet
}

func (ces componentEntitySet) intersect(others ...componentEntitySet) componentEntitySet {
	outSet := ces.clone()

	for e := range ces {
		for _, os := range others {
			if !os.contains(e) {
				outSet.remove(e)
			}
		}
	}

	return outSet
}

func (ces componentEntitySet) difference(others ...componentEntitySet) componentEntitySet {
	outSet := ces.clone()

	for e := range outSet {
		for _, os := range others {
			if os.contains(e) {
				outSet.remove(e)
			}
		}
	}

	return outSet
}

func union(sets ...componentEntitySet) componentEntitySet {
	if numSets := len(sets); numSets == 0 {
		return componentEntitySet{}
	} else if numSets == 1 {
		return sets[0]
	}

	return sets[0].union(sets[1:]...)
}

func intersect(sets ...componentEntitySet) componentEntitySet {
	if numSets := len(sets); numSets == 0 {
		return componentEntitySet{}
	} else if numSets == 1 {
		return sets[0]
	}

	// Sort from smallest to largest set; any intersection must be a subset of the smallest set,
	// so using that as our base set results in the fewest number of comparisons.
	sort.Slice(sets, func(i, j int) bool { return len(sets[i]) < len(sets[j]) })

	return sets[0].intersect(sets[1:]...)
}
