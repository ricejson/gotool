package queue

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayQueue_Offer(t *testing.T) {
	testCases := []struct {
		name       string
		capacity   int
		offerElems []string
		wantSize   int
		wantTop    string
		wantErr    error
	}{
		{
			name:       "多个元素入队（元素个数小于容量）",
			capacity:   4,
			offerElems: []string{"aa", "bb", "cc"},
			wantSize:   3,
			wantTop:    "aa",
			wantErr:    nil,
		},
		{
			name:       "多个元素入队（元素个数等于容量）",
			capacity:   4,
			offerElems: []string{"aa", "bb", "cc", "dd"},
			wantSize:   4,
			wantTop:    "aa",
			wantErr:    nil,
		},
		{
			name:       "多个元素入队（元素个数大于容量）",
			capacity:   4,
			offerElems: []string{"aa", "bb", "cc", "dd", "ee"},
			wantSize:   4,
			wantTop:    "aa",
			wantErr:    errors.New("queue is full"),
		},
		{
			name:       "多个元素入队（元素个数远远大于容量）",
			capacity:   4,
			offerElems: []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg"},
			wantSize:   4,
			wantTop:    "aa",
			wantErr:    errors.New("queue is full"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arrayQueue := NewArrayQueue[string](tc.capacity)
			var err error
			for _, elem := range tc.offerElems {
				err = arrayQueue.Offer(elem)
			}
			assert.Equal(t, err, tc.wantErr)
			top, err := arrayQueue.Peek()
			require.NoError(t, err)
			assert.Equal(t, top, tc.wantTop)
			assert.Equal(t, arrayQueue.Size(), tc.wantSize)
		})
	}
}

func TestArrayQueue_Poll(t *testing.T) {
	testCases := []struct {
		name       string
		capacity   int
		offerElems []string
		pollCount  int
		wantSize   int
		wantTop    string
		wantErr    error
	}{
		{
			name:       "多个元素出队（元素个数大于出队次数）",
			capacity:   4,
			offerElems: []string{"aa", "bb", "cc"},
			pollCount:  2,
			wantSize:   1,
			wantTop:    "cc",
			wantErr:    nil,
		},
		{
			name:       "多个元素出队（元素个数等于出队次数）",
			capacity:   4,
			offerElems: []string{"aa", "bb", "cc"},
			pollCount:  3,
			wantSize:   0,
			wantTop:    "",
			wantErr:    nil,
		},
		{
			name:       "多个元素出队（元素个数小于出队次数）",
			capacity:   4,
			offerElems: []string{"aa", "bb", "cc"},
			pollCount:  4,
			wantSize:   0,
			wantTop:    "",
			wantErr:    errors.New("queue is empty"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arrayQueue := NewArrayQueue[string](tc.capacity)
			var (
				top string
				err error
			)
			for _, elem := range tc.offerElems {
				err = arrayQueue.Offer(elem)
			}
			for i := 0; i < tc.pollCount; i++ {
				top, err = arrayQueue.Poll()
			}
			assert.Equal(t, err, tc.wantErr)
			top, err = arrayQueue.Peek()
			assert.Equal(t, top, tc.wantTop)
			assert.Equal(t, arrayQueue.Size(), tc.wantSize)
		})
	}
}

func TestArrayQueue_Integration(t *testing.T) {
	arrayQueue := NewArrayQueue[string](4)
	var (
		top string
		err error
	)
	err = arrayQueue.Offer("aa")
	assert.Equal(t, err, nil)
	err = arrayQueue.Offer("bb")
	assert.Equal(t, err, nil)
	err = arrayQueue.Offer("cc")
	assert.Equal(t, err, nil)
	err = arrayQueue.Offer("dd")
	assert.Equal(t, err, nil)
	err = arrayQueue.Offer("ee")
	assert.Equal(t, err, errors.New("queue is full"))
	top, err = arrayQueue.Poll()
	assert.Equal(t, top, "aa")
	assert.Equal(t, err, nil)
	top, err = arrayQueue.Poll()
	assert.Equal(t, top, "bb")
	assert.Equal(t, err, nil)
	top, err = arrayQueue.Poll()
	assert.Equal(t, top, "cc")
	assert.Equal(t, err, nil)
	top, err = arrayQueue.Poll()
	assert.Equal(t, top, "dd")
	assert.Equal(t, err, nil)
	top, err = arrayQueue.Poll()
	assert.Equal(t, top, "")
	assert.Equal(t, err, errors.New("queue is empty"))
	err = arrayQueue.Offer("ff")
	assert.Equal(t, err, nil)
	err = arrayQueue.Offer("gg")
	assert.Equal(t, err, nil)
	top, err = arrayQueue.Poll()
	assert.Equal(t, top, "ff")
	assert.Equal(t, err, nil)
	err = arrayQueue.Offer("hh")
	assert.Equal(t, err, nil)
	assert.Equal(t, err, nil)
	top, err = arrayQueue.Poll()
	assert.Equal(t, top, "gg")
	assert.Equal(t, err, nil)
	top, err = arrayQueue.Poll()
	assert.Equal(t, top, "hh")
	assert.Equal(t, err, nil)
	assert.Equal(t, arrayQueue.IsEmpty(), true)
}
