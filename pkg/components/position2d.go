package components

import "github.com/B1tVect0r/ymir/pkg/ecs"

func (pos *Position2D) Type() ecs.ComponentID {
	return ecs.ComponentID((&Position2D{}).ProtoReflect().Descriptor().FullName())
}
