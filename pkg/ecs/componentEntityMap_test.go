package ecs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCEMAdd(t *testing.T) {
	cID := ComponentID("c")
	eID := EntityHandle("e")

	cem := componentEntityMap{}

	err := cem.add(cID, eID)
	require.NoError(t, err)
	require.Equal(t, 1, len(cem))
	require.Equal(t, 1, len(cem[cID]))

	err = cem.add(cID, eID)
	require.Error(t, err)
}

func TestCEMRemove(t *testing.T) {
	cID := ComponentID("c")
	eID := EntityHandle("e")

	cem := componentEntityMap{cID: componentEntitySet{eID: struct{}{}}}

	cem.remove(cID, "other")
	require.Equal(t, 1, len(cem))
	require.Equal(t, 1, len(cem[cID]))

	cem.remove(cID, eID)
	require.Equal(t, 1, len(cem))
	require.Equal(t, 0, len(cem[cID]))
}
