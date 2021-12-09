package ecs

import (
	"math/rand"
	"testing"
	"time"

	"github.com/B1tVect0r/ymir/pkg/types/math/vector2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

type testComponent1 struct {
	vector2.Vector2
}

func (tc1 *testComponent1) Type() ComponentID {
	return "tc1"
}

type testComponent2 struct {
	vector2.Vector2
}

func (tc2 *testComponent2) Type() ComponentID {
	return "tc2"
}

func TestRegisterEntity(t *testing.T) {
	r := NewRegistry()

	c := &testComponent1{}
	cID := c.Type()
	e := &Entity{ID: uuid.NewString()}

	eID, err := r.RegisterEntity(e, c)
	require.NoError(t, err)

	entry, ok := r.em[eID]
	require.True(t, ok)
	if assert.Equal(t, len(entry.components), 1) {
		_, ok = entry.components[cID]
		require.True(t, ok)
	}

	es, ok := r.entitiesByComponent[cID]
	require.True(t, ok)
	if assert.Equal(t, len(es), 1) {
		_, ok = es[eID]
		require.True(t, ok)
	}
}

func TestMultiRoutineRegisterEntity(t *testing.T) {
	r := NewRegistry()

	eg := errgroup.Group{}

	e := &Entity{ID: uuid.NewString()}

	for i := 0; i < 5; i++ {
		eg.Go(func() error {

			time.Sleep(time.Duration(rand.Int31n(500)) * time.Millisecond)
			_, err := r.RegisterEntity(e)
			return err
		})
	}

	require.Error(t, eg.Wait())
}

func TestUnregisterEntity(t *testing.T) {
	r := NewRegistry()

	e := &Entity{ID: uuid.NewString()}
	c := &testComponent1{}
	eID, err := r.RegisterEntity(e, c)
	require.NoError(t, err)
	require.Equal(t, 1, len(r.em))
	require.Equal(t, 1, len(r.entitiesByComponent))
	require.Equal(t, 1, len(r.entitiesByComponent[c.Type()]))

	require.NoError(t, r.UnregisterEntity(eID))
	require.Equal(t, 0, len(r.em))
	require.Equal(t, 1, len(r.entitiesByComponent))
	require.Equal(t, 0, len(r.entitiesByComponent[c.Type()]))

}

func TestAddComponentsToEntity(t *testing.T) {
	r := NewRegistry()

	eID, err := r.RegisterEntity(&Entity{ID: uuid.NewString()})

	require.NoError(t, err)

	c1 := &testComponent1{}
	c2 := &testComponent2{}

	require.NoError(t, r.AddComponentsToEntity(eID, c1, c2))

	entry, ok := r.em[eID]
	require.True(t, ok)
	for _, cID := range []ComponentID{c1.Type(), c2.Type()} {
		_, ok = entry.components[cID]
		require.True(t, ok)

		es, ok := r.entitiesByComponent[cID]
		require.True(t, ok)

		_, ok = es[eID]
		require.True(t, ok)
	}

	require.Error(t, r.AddComponentsToEntity(eID, c1))
}

func TestRemoveComponentsFromEntity(t *testing.T) {
	r := NewRegistry()

	e := &Entity{ID: uuid.NewString()}
	c1 := &testComponent1{}
	c2 := &testComponent2{}
	eID, err := r.RegisterEntity(e, c1, c2)
	require.NoError(t, err)

	empty, err := r.RemoveComponentsFromEntity(eID, c1.Type())
	require.NoError(t, err)
	require.False(t, empty)
}

func TestGetComponentsFromEntity(t *testing.T) {
	r := NewRegistry()

	e := &Entity{ID: uuid.NewString()}
	c1 := &testComponent1{}
	c2 := &testComponent2{}
	eID, err := r.RegisterEntity(e, c1, c2)
	require.NoError(t, err)

	cs, err := r.GetComponentsFromEntity(eID, c1.Type())
	require.NoError(t, err)
	if assert.Equal(t, 1, len(cs)) {
		require.Equal(t, cs[0], c1)
	}

	cs, err = r.GetComponentsFromEntity(eID, c1.Type(), c2.Type())
	require.NoError(t, err)
	if assert.Equal(t, 2, len(cs)) {
		require.Equal(t, cs[0], c1)
		require.Equal(t, cs[1], c2)
	}

	_, err = r.GetComponentsFromEntity(eID, c1.Type(), c2.Type(), "dne")
	require.Error(t, err)
}

func TestQueryEntities(t *testing.T) {
	r := NewRegistry()

	c1 := &testComponent1{}
	c2 := &testComponent2{}

	eID, err := r.RegisterEntity(&Entity{ID: uuid.NewString()}, c1, c2)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		qm       SystemQueryMode
		incl     []ComponentID
		excl     []ComponentID
		expected int
	}{
		{
			"exclusive c1",
			Exclusive,
			[]ComponentID{c1.Type()},
			[]ComponentID{},
			0,
		},
		{
			"inclusive c1",
			Inclusive,
			[]ComponentID{c1.Type()},
			[]ComponentID{},
			1,
		},
		{
			"exclusive c1 and c2",
			Exclusive,
			[]ComponentID{c1.Type(), c2.Type()},
			[]ComponentID{},
			1,
		},
		{
			"exclusive c1, explicitly exclude c2",
			Exclusive,
			[]ComponentID{c1.Type()},
			[]ComponentID{c2.Type()},
			0,
		},
		{
			"inclusive c1, explicitly exclude c2",
			Inclusive,
			[]ComponentID{c1.Type()},
			[]ComponentID{c2.Type()},
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eIDs := r.QueryEntities(tc.qm, tc.incl, tc.excl)
			if numEntities := len(eIDs); assert.Equal(t, tc.expected, len(eIDs)) && numEntities > 0 {
				_, ok := eIDs[eID]
				require.True(t, ok)
			}
		})
	}
}
