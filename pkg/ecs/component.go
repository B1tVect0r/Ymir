package ecs

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ComponentID protoreflect.FullName

type Component interface {
	proto.Message
	Type() ComponentID
}
