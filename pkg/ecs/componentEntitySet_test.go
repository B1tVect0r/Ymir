package ecs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCESAdd(t *testing.T) {
	es := componentEntitySet{}
	require.Equal(t, 0, len(es))
	es.add("test")
	require.Equal(t, 1, len(es))

	// Adding an already-existing entry should do nothing
	es.add("test")
	require.Equal(t, 1, len(es))
}

func TestCESRemove(t *testing.T) {
	es := componentEntitySet{"test": struct{}{}}
	require.Equal(t, 1, len(es))

	// Removing a non-existent entry does nothing
	es.remove("dne")
	require.Equal(t, 1, len(es))

	es.remove("test")
	require.Equal(t, 0, len(es))
}

func TestCESContains(t *testing.T) {
	es := componentEntitySet{"test": struct{}{}}
	require.False(t, es.contains("dne"))
	require.True(t, es.contains("test"))
}

func TestCESClone(t *testing.T) {
	es := componentEntitySet{"test": struct{}{}}
	clone := es.clone()
	require.Equal(t, es, clone)

	// Modifying clone does not change es
	clone.add("other")
	require.NotEqual(t, es, clone)
}

func TestCESUnion(t *testing.T) {
	a := componentEntitySet{"a": struct{}{}}
	b := componentEntitySet{"b": struct{}{}}
	c := componentEntitySet{"c": struct{}{}}

	u := a.union(b, c)
	require.Equal(t, 3, len(u))
}

func TestCESIntersect(t *testing.T) {
	a := componentEntitySet{"a": struct{}{}}
	b := componentEntitySet{"b": struct{}{}}
	c := componentEntitySet{"c": struct{}{}}

	i := a.intersect(b, c)
	require.Equal(t, 0, len(i))

	a.add("common")
	b.add("common")
	c.add("common")

	i = a.intersect(b, c)
	require.Equal(t, 1, len(i))
}

func TestCESDifference(t *testing.T) {
	a := componentEntitySet{"a": struct{}{}}
	b := componentEntitySet{"b": struct{}{}}
	c := componentEntitySet{"c": struct{}{}}

	d := a.difference(b, c)
	require.Equal(t, 1, len(d))

	b.add("a")

	d = a.difference(b, c)
	require.Equal(t, 0, len(d))
}
