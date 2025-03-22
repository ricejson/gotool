package stack

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListStack_Integration(t *testing.T) {
	listStack := NewListStack[string]()
	err := listStack.Push("aa")
	require.NoError(t, err)
	err = listStack.Push("bb")
	require.NoError(t, err)
	err = listStack.Push("cc")
	require.NoError(t, err)
	err = listStack.Push("dd")
	require.NoError(t, err)
	peek, err := listStack.Peek()
	require.NoError(t, err)
	assert.Equal(t, peek, "dd")
	pop, err := listStack.Pop()
	require.NoError(t, err)
	assert.Equal(t, pop, "dd")
	pop, err = listStack.Pop()
	pop, err = listStack.Pop()
	topVal, err := listStack.Peek()
	require.NoError(t, err)
	assert.Equal(t, topVal, "aa")
	pop, err = listStack.Pop()
	topVal, err = listStack.Peek()
	assert.Equal(t, err, errors.New("stack is empty"))
}

func TestListStack_Push(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
