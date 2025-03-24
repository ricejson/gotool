package bitmap

import (
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBitMap_Integration(t *testing.T) {
	bitMap := NewBitMap()
	err := bitMap.Set(1)
	require.NoError(t, err)
	err = bitMap.Set(3)
	require.NoError(t, err)
	err = bitMap.Set(4)
	require.NoError(t, err)
	exists, err := bitMap.Contains(7)
	assert.Equal(t, exists, false)
	require.NoError(t, err)
	err = bitMap.Set(7)
	require.NoError(t, err)
	exists, err = bitMap.Contains(7)
	assert.Equal(t, exists, true)
	require.NoError(t, err)
	err = bitMap.Reset(7)
	require.NoError(t, err)
	exists, err = bitMap.Contains(7)
	assert.Equal(t, exists, false)
	require.NoError(t, err)
	assert.Equal(t, bitMap.Size(), 3)
}
