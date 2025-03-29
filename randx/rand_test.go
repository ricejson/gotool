package randx

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRand(t *testing.T) {
	randVal, err := Int(-1, 10)
	assert.Equal(t, err, errors.New("minInclude < 0"))
	assert.Equal(t, randVal, 0)
	randVal, err = Int(12, 10)
	assert.Equal(t, err, errors.New("minInclude >= maxExclude"))
	assert.Equal(t, randVal, 0)
	randVal, err = Int(10, 100)
	require.NoError(t, err)
	require.True(t, randVal >= 10 && randVal < 100)
}
