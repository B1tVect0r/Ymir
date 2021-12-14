package orbiter

import (
	"log"
	"time"

	"github.com/B1tVect0r/ymir/pkg/components/orbit"
	"github.com/B1tVect0r/ymir/pkg/components/pos2d"
	"github.com/B1tVect0r/ymir/pkg/ecs"
)

const SystemID ecs.SystemID = 1

var _ ecs.System = (*OrbiterSystem)(nil)

type OrbiterSystem struct {
	elapsed time.Duration
}

func (os *OrbiterSystem) ID() ecs.SystemID {
	return SystemID
}

func (os *OrbiterSystem) QueryMode() ecs.SystemQueryMode {
	return ecs.Exclusive
}

func (os *OrbiterSystem) IncludeTypes() []ecs.ComponentID {
	return []ecs.ComponentID{
		orbit.ComponentID,
		pos2d.ComponentID,
	}
}

func (os *OrbiterSystem) ExcludeTypes() []ecs.ComponentID {
	return []ecs.ComponentID{}
}

func (os *OrbiterSystem) Process(dt time.Duration, operativeSets ecs.OperativeSets) {
	os.elapsed += dt

	for entity, components := range operativeSets {
		orbit := components[orbit.ComponentID].(*orbit.T)
		pos := components[pos2d.ComponentID].(*pos2d.T)
		pos.Loc = orbit.PositionAt(orbit.AngleAt(os.elapsed))

		log.Printf("Entity %d position: %v", entity, pos)
	}
}
