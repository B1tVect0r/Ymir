package ecs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCESAdd(t *testing.T) {
	es := componentEntitySet{}
	require.Equal(t, 0, len(es))
	es.add(1)
	require.Equal(t, 1, len(es))

	// Adding an already-existing entry should do nothing
	es.add(1)
	require.Equal(t, 1, len(es))
}

func TestCESRemove(t *testing.T) {
	es := componentEntitySet{1: struct{}{}}
	require.Equal(t, 1, len(es))

	// Removing a non-existent entry does nothing
	es.remove(2)
	require.Equal(t, 1, len(es))

	es.remove(1)
	require.Equal(t, 0, len(es))
}

func TestCESContains(t *testing.T) {
	es := componentEntitySet{1: struct{}{}}
	require.False(t, es.contains(2))
	require.True(t, es.contains(1))
}

func TestCESClone(t *testing.T) {
	es := componentEntitySet{1: struct{}{}}
	clone := es.clone()
	require.Equal(t, es, clone)

	// Modifying clone does not change es
	clone.add(2)
	require.NotEqual(t, es, clone)
}

func TestCESUnion(t *testing.T) {
	a := componentEntitySet{1: struct{}{}}
	b := componentEntitySet{2: struct{}{}}
	c := componentEntitySet{3: struct{}{}}

	u := a.union(b, c)
	require.Equal(t, 3, len(u))
}

func TestCESIntersect(t *testing.T) {
	a := componentEntitySet{1: struct{}{}}
	b := componentEntitySet{2: struct{}{}}
	c := componentEntitySet{3: struct{}{}}

	i := a.intersect(b, c)
	require.Equal(t, 0, len(i))

	a.add(4)
	b.add(4)
	c.add(4)

	i = a.intersect(b, c)
	require.Equal(t, 1, len(i))
}

func TestCESDifference(t *testing.T) {
	a := componentEntitySet{1: struct{}{}}
	b := componentEntitySet{2: struct{}{}}
	c := componentEntitySet{3: struct{}{}}

	d := a.difference(b, c)
	require.Equal(t, 1, len(d))

	b.add(1)

	d = a.difference(b, c)
	require.Equal(t, 0, len(d))
}
