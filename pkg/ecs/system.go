package ecs

import (
	"time"

	"log"
)

type OperativeSets map[EntityID]map[ComponentID]Component

type SystemQueryMode int

const (
	Exclusive SystemQueryMode = iota
	Inclusive
)

type SystemID uint16

type System interface {
	ID() SystemID
	QueryMode() SystemQueryMode
	IncludeTypes() []ComponentID
	ExcludeTypes() []ComponentID
	Process(dt time.Duration, operativeSets OperativeSets)
}

type systemHarness struct {
	s System
}

func (sw *systemHarness) id() SystemID { return sw.s.ID() }

func (sw *systemHarness) process(r *Registry, dt time.Duration) error {
	entryTime := time.Now().UTC()
	incl := sw.s.IncludeTypes()
	es := r.QueryEntities(sw.s.QueryMode(), incl, sw.s.ExcludeTypes())

	if len(es) == 0 {
		return nil
	}

	operativeSets := make(OperativeSets, len(es))

	for e := range es {
		cs, err := r.GetComponentsFromEntity(e, incl...)
		if err != nil {
			return err
		}
		operativeSets[e] = cs
	}

	sw.s.Process(dt+time.Since(entryTime), operativeSets)
	return nil
}

func Process(r *Registry, dt time.Duration) {
	entry := time.Now().UTC()

	for _, sys := range r.systems {
		if err := sys.process(r, dt+time.Since(entry)); err != nil {
			log.Printf("System failed to process: %s", err.Error())
		}
	}
}

func RegisterSystems(r *Registry, systems ...System) {
	wrappers := make([]*systemHarness, len(systems))
	for i, s := range systems {
		wrappers[i] = &systemHarness{s: s}
	}

	r.registerSystems(wrappers...)
}

func UnregisterSystems(r *Registry, sIDs ...SystemID) {
	r.unregisterSystems(sIDs...)
}
