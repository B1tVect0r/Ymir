package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/B1tVect0r/ymir/pkg/components"
	"github.com/B1tVect0r/ymir/pkg/ecs"
	"github.com/B1tVect0r/ymir/pkg/types/math/vector2"
	"github.com/B1tVect0r/ymir/pkg/world"
)

type mySystem struct{}

func (s *mySystem) ID() ecs.SystemID {
	return "SystemA"
}

func (s *mySystem) QueryMode() ecs.SystemQueryMode {
	return ecs.Exclusive
}

func (s *mySystem) IncludeTypes() []ecs.ComponentID {
	return []ecs.ComponentID{
		(&components.Position2D{}).Type(),
	}
}

func (s *mySystem) ExcludeTypes() []ecs.ComponentID {
	return []ecs.ComponentID{}
}

func (s *mySystem) Process(dt time.Duration, operativeSets ecs.OperativeSets) {
	stats := make([]string, 0, len(operativeSets))
	for e, cs := range operativeSets {
		stats = append(stats, fmt.Sprintf("%s: %d components", e, len(cs)))
	}

	fmt.Fprintf(os.Stdout, "System A Stats: \n\t%s\n", strings.Join(stats, "\n\t"))
}

type myOtherSystem struct{}

func (s *myOtherSystem) ID() ecs.SystemID {
	return "SystemB"
}

func (s *myOtherSystem) QueryMode() ecs.SystemQueryMode {
	return ecs.Exclusive
}

func (s *myOtherSystem) IncludeTypes() []ecs.ComponentID {
	return []ecs.ComponentID{
		(&components.Position2D{}).Type(),
		(&components.Facing{}).Type(),
	}
}

func (s *myOtherSystem) ExcludeTypes() []ecs.ComponentID {
	return []ecs.ComponentID{}
}

func (s *myOtherSystem) Process(dt time.Duration, operativeSets ecs.OperativeSets) {
	stats := make([]string, 0, len(operativeSets))
	for e, cs := range operativeSets {
		stats = append(stats, fmt.Sprintf("%s: %d components", e, len(cs)))
	}

	fmt.Fprintf(os.Stdout, "System B Stats: \n\t%s\n", strings.Join(stats, "\n\t"))
}

func main() {
	w := world.NewWorld()
	e1, _ := w.CreateEntity(&components.Position2D{Position: vector2.Right()})
	e2, _ := w.CreateEntity(&components.Position2D{Position: vector2.Up()})

	fmt.Fprintf(os.Stdout, "e1: %s\n", e1)
	fmt.Fprintf(os.Stdout, "e2: %s\n", e2)

	t1 := time.AfterFunc(1*time.Second, func() {
		fmt.Fprintf(os.Stdout, "Adding facing component to E2 (%s)\n", e2)
		w.AddComponentsToEntity(e2, &components.Facing{})
	})
	t2 := time.AfterFunc(2*time.Second, func() {
		fmt.Fprintf(os.Stdout, "Removing facing component from E2 (%s)\n", e2)
		w.RemoveComponentsFromEntity(e2, (&components.Facing{}).Type())
	})
	t3 := time.AfterFunc(3*time.Second, func() {
		fmt.Fprintf(os.Stdout, "Deleting entity E2 (%s)\n", e2)
		w.DestroyEntity(e2)
	})

	defer t1.Stop()
	defer t2.Stop()
	defer t3.Stop()

	w.RegisterSystems(&mySystem{}, &myOtherSystem{})

	done := make(chan struct{})

	go w.Start(done)

	time.Sleep(5 * time.Second)
	done <- struct{}{}
}
