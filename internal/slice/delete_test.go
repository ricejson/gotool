package slice

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name        string
		src         []int
		index       int
		wantSlice   []int
		wantElement int
		wantErr     error
	}{
		{
			name:        "索引越界",
			src:         []int{1, 2, 3},
			index:       -1,
			wantSlice:   nil,
			wantElement: 0,
			wantErr:     errors.New("index out of range"),
		},
		{
			name:        "索引越界",
			src:         []int{1, 2, 3},
			index:       3,
			wantSlice:   nil,
			wantElement: 0,
			wantErr:     errors.New("index out of range"),
		},
		{
			name:        "删除边界元素成功",
			src:         []int{1, 2, 3},
			index:       2,
			wantSlice:   []int{1, 2},
			wantElement: 3,
			wantErr:     nil,
		},
		{
			name:        "删除中间元素成功",
			src:         []int{1, 2, 3},
			index:       1,
			wantSlice:   []int{1, 3},
			wantElement: 2,
			wantErr:     nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			src := tc.src
			newSlice, element, err := Delete(src, tc.index)
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, newSlice, tc.wantSlice)
			assert.Equal(t, element, tc.wantElement)

		})
	}
}
