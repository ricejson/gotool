package stack

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestArrayStack_Integration(t *testing.T) {
	arrayStack := NewArrayStack[string](4)
	err := arrayStack.Push("aa")
	require.NoError(t, err)
	err = arrayStack.Push("bb")
	require.NoError(t, err)
	err = arrayStack.Push("cc")
	require.NoError(t, err)
	err = arrayStack.Push("dd")
	require.NoError(t, err)
	assert.Equal(t, arrayStack, ArrayStack[string]{
		elems:    []string{"aa", "bb", "cc", "dd"},
		top:      4,
		capacity: 4,
	})
	err = arrayStack.Push("ee")
	require.NoError(t, err)
	assert.Equal(t, arrayStack, ArrayStack[string]{
		elems:    []string{"aa", "bb", "cc", "dd", "ee", "", "", ""},
		top:      5,
		capacity: 8,
	})
	peek, err := arrayStack.Peek()
	require.NoError(t, err)
	assert.Equal(t, peek, "ee")
	pop, err := arrayStack.Pop()
	require.NoError(t, err)
	assert.Equal(t, pop, "ee")
	assert.Equal(t, arrayStack, ArrayStack[string]{
		elems:    []string{"aa", "bb", "cc", "dd", "ee", "", "", ""},
		top:      4,
		capacity: 8,
	})
	pop, err = arrayStack.Pop()
	pop, err = arrayStack.Pop()
	pop, err = arrayStack.Pop()
	pop, err = arrayStack.Pop()
	pop, err = arrayStack.Peek()
	assert.Equal(t, err, errors.New("stack is empty"))
	pop, err = arrayStack.Pop()
	assert.Equal(t, err, errors.New("stack is empty"))
}

func TestArrayStack_Push(t *testing.T) {
	testCases := []struct {
		name     string
		inputArr []string
		wantRes  Stack[string]
	}{
		{
			name:     "测试Push方法",
			inputArr: []string{},
			wantRes:  NewArrayStack[string](4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arrayStack := NewArrayStack[string](4)
			for _, v := range tc.inputArr {
				err := arrayStack.Push(v)
				require.NoError(t, err)
			}
			assert.Equal(t, tc.wantRes, arrayStack)
		})
	}
}
