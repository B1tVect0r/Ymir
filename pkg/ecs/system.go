package ecs

import (
	"time"

	"log"
)

type OperativeSets map[EntityHandle][]Component

type SystemQueryMode int

const (
	Exclusive SystemQueryMode = iota
	Inclusive
)

type SystemID string

type System interface {
	ID() SystemID
	QueryMode() SystemQueryMode
	IncludeTypes() []ComponentID
	ExcludeTypes() []ComponentID
	Process(dt time.Duration, operativeSets OperativeSets)
}

type systemWrapper struct {
	del System
}

func (sw *systemWrapper) id() SystemID { return sw.del.ID() }

func (sw *systemWrapper) process(r *Registry, dt time.Duration) error {
	entryTime := time.Now().UTC()
	incl := sw.del.IncludeTypes()
	es := r.QueryEntities(sw.del.QueryMode(), incl, sw.del.ExcludeTypes())

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

	sw.del.Process(dt+time.Since(entryTime), operativeSets)
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
	wrappers := make([]*systemWrapper, len(systems))
	for i, s := range systems {
		wrappers[i] = &systemWrapper{del: s}
	}

	r.registerSystems(wrappers...)
}

func UnregisterSystems(r *Registry, sIDs ...SystemID) {
	r.unregisterSystems(sIDs...)
}
