package world

import (
	"testing"
	"time"

	"github.com/B1tVect0r/ymir/pkg/ecs"
)

type worldBenchmarkComponent struct {
	id ecs.ComponentID
}

func (wbc *worldBenchmarkComponent) ID() ecs.ComponentID { return wbc.id }

type worldBenchmarkSystem struct {
	include []ecs.ComponentID
}

func (wbs *worldBenchmarkSystem) ID() ecs.SystemID                               { return 1 }
func (wbs *worldBenchmarkSystem) QueryMode() ecs.SystemQueryMode                 { return ecs.Exclusive }
func (wbs *worldBenchmarkSystem) IncludeTypes() []ecs.ComponentID                { return wbs.include }
func (wbs *worldBenchmarkSystem) ExcludeTypes() []ecs.ComponentID                { return []ecs.ComponentID{} }
func (wbs *worldBenchmarkSystem) Process(dt time.Duration, es ecs.OperativeSets) {}

func BenchmarkRunningWorld(b *testing.B) {
	w := NewWorld()

	components := make([]ecs.Component, 64)
	for i := range components {
		components[i] = &worldBenchmarkComponent{id: ecs.ComponentID(i + 1)}
		if i > 0 {
			w.CreateEntity(components[0], components[1:i]...)
		}
	}

	w.RegisterSystems(&worldBenchmarkSystem{})

	t := time.Now()
	for i := 0; i < b.N; i++ {
		nt := time.Now()
		dt := nt.Sub(t)
		t = nt

		ecs.Process(w.reg, dt)
	}
}
