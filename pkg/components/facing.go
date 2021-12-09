package components

import "github.com/B1tVect0r/ymir/pkg/ecs"

func (f *Facing) Type() ecs.ComponentID {
	return ecs.ComponentID((&Facing{}).ProtoReflect().Descriptor().FullName())
}
