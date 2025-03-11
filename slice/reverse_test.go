package slice

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name    string
		src     []string
		wantRes []string
	}{
		{
			name:    "切片为nil",
			src:     nil,
			wantRes: []string{},
		},
		{
			name:    "切片为空",
			src:     []string{},
			wantRes: []string{},
		},
		{
			name:    "切片只有一个元素",
			src:     []string{"aa"},
			wantRes: []string{"aa"},
		},
		{
			name:    "切片有多个奇数元素",
			src:     []string{"aa", "bb", "cc"},
			wantRes: []string{"cc", "bb", "aa"},
		},
		{
			name:    "切片有多个偶数元素",
			src:     []string{"aa", "bb", "cc", "dd"},
			wantRes: []string{"dd", "cc", "bb", "aa"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			assert.Equal(t, tc.wantRes, Reverse[string](src))
		})
	}
}

func TestReverseSelf(t *testing.T) {
	testCases := []struct {
		name    string
		src     []string
		wantRes []string
	}{
		{
			name:    "切片为nil",
			src:     nil,
			wantRes: nil,
		},
		{
			name:    "切片为空",
			src:     []string{},
			wantRes: []string{},
		},
		{
			name:    "切片只有一个元素",
			src:     []string{"aa"},
			wantRes: []string{"aa"},
		},
		{
			name:    "切片有多个奇数元素",
			src:     []string{"aa", "bb", "cc"},
			wantRes: []string{"cc", "bb", "aa"},
		},
		{
			name:    "切片有多个偶数元素",
			src:     []string{"aa", "bb", "cc", "dd"},
			wantRes: []string{"dd", "cc", "bb", "aa"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			ReverseSelf[string](src)
			assert.Equal(t, tc.wantRes, src)
		})
	}
}
