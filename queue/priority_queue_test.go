package queue

import (
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPriorityQueue_Integration(t *testing.T) {
	pq := NewPriorityQueue[int](10, func(e1, e2 int) int { return e1 - e2 })
	err := pq.Offer(4)
	require.NoError(t, err)
	err = pq.Offer(8)
	require.NoError(t, err)
	err = pq.Offer(2)
	require.NoError(t, err)
	err = pq.Offer(10)
	require.NoError(t, err)
	err = pq.Offer(1)
	require.NoError(t, err)
	peek, err := pq.Poll()
	require.NoError(t, err)
	assert.Equal(t, peek, 1)
	peek, err = pq.Poll()
	require.NoError(t, err)
	assert.Equal(t, peek, 2)
	peek, err = pq.Poll()
	require.NoError(t, err)
	assert.Equal(t, peek, 4)
	peek, err = pq.Poll()
	require.NoError(t, err)
	assert.Equal(t, peek, 8)
	peek, err = pq.Poll()
	require.NoError(t, err)
	assert.Equal(t, peek, 10)
	assert.Equal(t, pq.IsEmpty(), true)
}
