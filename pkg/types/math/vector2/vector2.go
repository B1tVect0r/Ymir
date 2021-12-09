package vector2

import (
	"math"
)

func Right() *Vector2 { return &Vector2{X: 1} }
func Up() *Vector2    { return &Vector2{Y: 1} }

func (v *Vector2) Add(other *Vector2) *Vector2 {
	return &Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v *Vector2) Distance(other *Vector2) float64 {
	return math.Sqrt(v.DistanceSquared(other))
}

func (v *Vector2) DistanceSquared(other *Vector2) float64 {
	dx := other.X - v.X
	dy := other.Y - v.Y

	return dx*dx + dy*dy
}
