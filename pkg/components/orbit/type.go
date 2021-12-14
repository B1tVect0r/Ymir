package orbit

import (
	"math"
	"time"

	"github.com/B1tVect0r/ymir/pkg/ecs"
	"github.com/B1tVect0r/ymir/pkg/types/math/vec2"
)

const ComponentID ecs.ComponentID = 42

type T struct {
	// Center of the orbit
	C vec2.T
	// Eccentricity; 0 <= e < 1
	Ecc float64
	// Semi-Major axis
	Major  float64
	minor  float64
	period time.Duration
}

func (o *T) ID() ecs.ComponentID {
	return ComponentID
}

func (o *T) Period() time.Duration {
	return o.period
}

func (o *T) AngleAt(t time.Duration) float64 {
	// Todo: Implement for orbits with Ecc != 0
	nt := t
	for ; nt.Seconds() > o.period.Seconds(); nt = nt - o.period {
	}

	return (nt.Seconds() / o.period.Seconds()) * 2 * math.Pi
}

func (o *T) PositionAt(rad float64) vec2.T {
	sinRad := math.Sin(rad)
	cosRad := math.Cos(rad)
	r := (o.Major * o.minor) / math.Sqrt((o.Major*o.Major*sinRad*sinRad)+(o.minor*o.minor*cosRad*cosRad))
	return o.C.Add(vec2.T{cosRad, sinRad}.Scale(r))
}

func NewOrbit(c vec2.T, ecc, maj, sgp float64) *T {
	return &T{
		C:      c,
		Ecc:    0,
		Major:  maj,
		minor:  maj * math.Sqrt(1-(ecc*ecc)),
		period: time.Duration(2*math.Pi*math.Sqrt(math.Pow(maj, 3)/sgp)) * time.Second,
	}
}
