package world

import "github.com/B1tVect0r/ymir/pkg/ecs"

func (w *World) CreateEntity(component ecs.Component, components ...ecs.Component) (ecs.EntityHandle, error) {
	return ecs.NewEntity(w.reg, component, components...)
}

func (w *World) DestroyEntity(eID ecs.EntityHandle) error {
	return ecs.Destroy(w.reg, eID)
}

func (w *World) AddComponentsToEntity(eID ecs.EntityHandle, components ...ecs.Component) error {
	return ecs.AddComponents(w.reg, eID, components...)
}

func (w *World) RemoveComponentsFromEntity(eID ecs.EntityHandle, cIDs ...ecs.ComponentID) error {
	return ecs.RemoveComponents(w.reg, eID, cIDs...)
}

func (w *World) RegisterSystems(system ecs.System, systems ...ecs.System) {
	ecs.RegisterSystems(w.reg, append([]ecs.System{system}, systems...)...)
}

func (w *World) UnregisterSystems(sID ecs.SystemID, sIDs ...ecs.SystemID) {
	ecs.UnregisterSystems(w.reg, append([]ecs.SystemID{sID}, sIDs...)...)
}
