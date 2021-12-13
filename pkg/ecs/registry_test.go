package ecs

import (
	"testing"

	"github.com/B1tVect0r/ymir/pkg/types/math/vec2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testComponent1 struct {
	vec2.T
}

func (tc1 *testComponent1) ID() ComponentID {
	return 1
}

type testComponent2 struct {
	vec2.T
}

func (tc2 *testComponent2) ID() ComponentID {
	return 2
}

func TestRegisterEntity(t *testing.T) {
	r := NewRegistry()

	c := &testComponent1{}
	cID := c.ID()
	eID, err := r.RegisterEntity(c)
	require.NoError(t, err)

	cm, ok := r.em[eID]
	require.True(t, ok)
	if assert.Equal(t, len(cm), 1) {
		_, ok = cm[cID]
		require.True(t, ok)
	}

	es, ok := r.entitiesByComponent[cID]
	require.True(t, ok)
	if assert.Equal(t, len(es), 1) {
		_, ok = es[eID]
		require.True(t, ok)
	}
}

func TestUnregisterEntity(t *testing.T) {
	r := NewRegistry()

	c := &testComponent1{}
	eID, err := r.RegisterEntity(c)
	require.NoError(t, err)
	require.Equal(t, 1, len(r.em))
	require.Equal(t, 1, len(r.entitiesByComponent))
	require.Equal(t, 1, len(r.entitiesByComponent[c.ID()]))

	require.NoError(t, r.UnregisterEntity(eID))
	require.Equal(t, 0, len(r.em))
	require.Equal(t, 1, len(r.entitiesByComponent))
	require.Equal(t, 0, len(r.entitiesByComponent[c.ID()]))

}

func TestAddComponentsToEntity(t *testing.T) {
	r := NewRegistry()

	c1 := &testComponent1{}
	eID, err := r.RegisterEntity(c1)

	require.NoError(t, err)

	c2 := &testComponent2{}

	require.NoError(t, r.AddComponentsToEntity(eID, c2))

	cm, ok := r.em[eID]
	require.True(t, ok)
	for _, cID := range []ComponentID{c1.ID(), c2.ID()} {
		_, ok = cm[cID]
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

	c1 := &testComponent1{}
	c2 := &testComponent2{}
	eID, err := r.RegisterEntity(c1, c2)
	require.NoError(t, err)

	empty, err := r.RemoveComponentsFromEntity(eID, c1.ID())
	require.NoError(t, err)
	require.False(t, empty)
}

func TestGetComponentsFromEntity(t *testing.T) {
	r := NewRegistry()

	c1 := &testComponent1{}
	c2 := &testComponent2{}
	eID, err := r.RegisterEntity(c1, c2)
	require.NoError(t, err)

	cs, err := r.GetComponentsFromEntity(eID, c1.ID())
	require.NoError(t, err)
	if assert.Equal(t, 1, len(cs)) {
		require.Equal(t, cs[c1.ID()], c1)
	}

	cs, err = r.GetComponentsFromEntity(eID, c1.ID(), c2.ID())
	require.NoError(t, err)
	if assert.Equal(t, 2, len(cs)) {
		require.Equal(t, cs[c1.ID()], c1)
		require.Equal(t, cs[c2.ID()], c2)
	}

	_, err = r.GetComponentsFromEntity(eID, c1.ID(), c2.ID(), 3)
	require.Error(t, err)
}

func TestQueryEntities(t *testing.T) {
	r := NewRegistry()

	c1 := &testComponent1{}
	c2 := &testComponent2{}

	eID, err := r.RegisterEntity(c1, c2)
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
			[]ComponentID{c1.ID()},
			[]ComponentID{},
			0,
		},
		{
			"inclusive c1",
			Inclusive,
			[]ComponentID{c1.ID()},
			[]ComponentID{},
			1,
		},
		{
			"exclusive c1 and c2",
			Exclusive,
			[]ComponentID{c1.ID(), c2.ID()},
			[]ComponentID{},
			1,
		},
		{
			"exclusive c1, explicitly exclude c2",
			Exclusive,
			[]ComponentID{c1.ID()},
			[]ComponentID{c2.ID()},
			0,
		},
		{
			"inclusive c1, explicitly exclude c2",
			Inclusive,
			[]ComponentID{c1.ID()},
			[]ComponentID{c2.ID()},
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

func BenchmarkQueryEntities(b *testing.B) {
	r := NewRegistry()

	for i := 0; i < 512; i++ {
		eID, err := r.RegisterEntity(&testComponent1{})
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		if i%2 == 0 {
			r.AddComponentsToEntity(eID, &testComponent2{})
		}

	}

	for i := 0; i < b.N; i++ {
		r.QueryEntities(Exclusive, []ComponentID{1}, []ComponentID{})
		r.QueryEntities(Inclusive, []ComponentID{1}, []ComponentID{})
		r.QueryEntities(Exclusive, []ComponentID{1, 2}, []ComponentID{})
		r.QueryEntities(Inclusive, []ComponentID{2}, []ComponentID{1})
	}
}
