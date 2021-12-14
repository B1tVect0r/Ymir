package orbit

import (
	"math"
	"testing"

	"github.com/B1tVect0r/ymir/pkg/types/math/vec2"
	"github.com/stretchr/testify/require"
)

func TestNewOrbit(t *testing.T) {
	o := NewOrbit(vec2.Zero, 0, 1, 1)
	require.Equal(t, o.Major, o.minor)
	require.True(t, vec2.Right.Equals(o.PositionAt(0)))
	require.True(t, vec2.Up.Equals(o.PositionAt(math.Pi/2)))
	require.True(t, vec2.Right.Scale(-1).Equals(o.PositionAt(math.Pi)))
	require.True(t, vec2.Up.Scale(-1).Equals(o.PositionAt(3*math.Pi/2)))

	epsilon := math.Nextafter(1, 2) - 1

	require.True(t, o.AngleAt(0) < epsilon)
	require.True(t, o.AngleAt(o.period/4)-math.Pi/2 < epsilon)
	require.True(t, o.AngleAt(o.period/2)-math.Pi < epsilon)
	require.True(t, o.AngleAt(3*o.period/4)-3*math.Pi/2 < epsilon)
	require.True(t, o.AngleAt(0)-o.AngleAt(o.period) < epsilon)
	require.True(t, o.AngleAt(5*o.period/4)-o.AngleAt(o.period/4) < epsilon)
}
