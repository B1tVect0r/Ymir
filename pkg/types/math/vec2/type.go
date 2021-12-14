package vec2

import (
	"math"
)

var epsilon = math.Nextafter(1, 2) - 1

type T [2]float64

var Zero = T{0, 0}
var Right = T{1, 0}
var Up = T{0, 1}

func (vec T) Scale(scale float64) T {
	vec[0] *= scale
	vec[1] *= scale
	return vec
}

func (vec T) ScaleX(scale float64) T {
	vec[0] *= scale
	return vec
}

func (vec T) ScaleY(scale float64) T {
	vec[1] *= scale
	return vec
}

func (vec T) Add(other T) T {
	vec[0] += other[0]
	vec[1] += other[1]
	return vec
}

func (vec T) Sub(other T) T {
	return vec.Add(other.Scale(-1))
}

func (vec T) Length() float64 {
	return math.Sqrt(vec.LengthSquared())
}

func (vec T) LengthSquared() float64 {
	return vec[0]*vec[0] + vec[1]*vec[1]
}

func (vec T) Normalize() T {
	return vec.Scale(1 / vec.Length())
}

func (vec T) Equals(other T) bool {
	return (other[0]-vec[0]) < epsilon &&
		(other[1]-vec[1]) < epsilon
}
