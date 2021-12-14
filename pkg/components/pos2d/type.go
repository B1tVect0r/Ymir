package pos2d

import (
	"github.com/B1tVect0r/ymir/pkg/ecs"
	"github.com/B1tVect0r/ymir/pkg/types/math/vec2"
)

const ComponentID ecs.ComponentID = 1

type T struct {
	Loc vec2.T
	// Rotation in radians
	Rot float64
}

func (pos *T) ID() ecs.ComponentID {
	return ComponentID
}
