package slice

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		src       []int
		index     int
		val       int
		wantErr   error
		wantSlice []int
	}{
		{
			name:      "索引越界",
			src:       []int{1, 2, 3},
			index:     -1,
			val:       5,
			wantErr:   errors.New("index out of range"),
			wantSlice: nil,
		},
		{
			name:      "索引越界",
			src:       []int{1, 2, 3},
			index:     4,
			val:       5,
			wantErr:   errors.New("index out of range"),
			wantSlice: nil,
		},
		{
			name:      "插入成功",
			src:       []int{1, 2, 3},
			index:     3,
			val:       4,
			wantErr:   nil,
			wantSlice: []int{1, 2, 3, 4},
		},
		{
			name:      "插入成功",
			src:       []int{1, 3, 4},
			index:     1,
			val:       2,
			wantErr:   nil,
			wantSlice: []int{1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			src := tc.src
			newSlice, err := Add(src, tc.index, tc.val)
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, newSlice, tc.wantSlice)
		})
	}
}
