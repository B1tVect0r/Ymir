package world

import (
	"time"

	"github.com/B1tVect0r/ymir/pkg/ecs"
)

type World struct {
	reg       *ecs.Registry
	targetFPS int
}

func (w *World) Start(done <-chan struct{}) {
	timer := time.NewTicker(time.Duration(1000/w.targetFPS) * time.Millisecond)
	t := time.Now()
	for {
		select {
		case <-done:
			timer.Stop()
			return
		case nt := <-timer.C:
			dt := nt.Sub(t)
			t = nt
			ecs.Process(w.reg, dt)
		}
	}
}

type worldConfig struct {
	ecsConfig ecs.RegistryConfig
	targetFPS int
}

func defaultWorldConfig() *worldConfig {
	return &worldConfig{
		targetFPS: 65,
	}
}

func (wc *worldConfig) applyOptions(opts ...WorldOption) {
	for _, o := range opts {
		o(wc)
	}
}

type WorldOption func(wc *worldConfig)

func NewWorld(opts ...WorldOption) *World {
	c := defaultWorldConfig()
	c.applyOptions(opts...)

	return &World{
		reg:       ecs.NewRegistryWithConfig(c.ecsConfig),
		targetFPS: c.targetFPS,
	}
}
