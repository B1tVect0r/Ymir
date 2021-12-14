package orbiter

import (
	"time"

	"github.com/B1tVect0r/ymir/pkg/components/orbit"
	"github.com/B1tVect0r/ymir/pkg/components/pos2d"
	"github.com/B1tVect0r/ymir/pkg/ecs"
)

const SystemID ecs.SystemID = 1

var _ ecs.System = (*System)(nil)

type System struct {
	elapsed time.Duration
}

func (s *System) ID() ecs.SystemID {
	return SystemID
}

func (s *System) QueryMode() ecs.SystemQueryMode {
	return ecs.Exclusive
}

func (s *System) IncludeTypes() []ecs.ComponentID {
	return []ecs.ComponentID{
		orbit.ComponentID,
		pos2d.ComponentID,
	}
}

func (s *System) ExcludeTypes() []ecs.ComponentID {
	return []ecs.ComponentID{}
}

func (s *System) Process(dt time.Duration, operativeSets ecs.OperativeSets) {
	s.elapsed += dt

	for _, components := range operativeSets {
		orbit := components[orbit.ComponentID].(*orbit.T)
		pos := components[pos2d.ComponentID].(*pos2d.T)
		pos.Loc = orbit.PositionAt(orbit.AngleAt(s.elapsed))
	}
}
