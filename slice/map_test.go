package slice

import (
	"github.com/go-playground/assert/v2"
	"strconv"
	"strings"
	"testing"
)

func TestSlice_Map(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		wantRes []string
	}{
		{
			name:    "切片为空",
			src:     []int{},
			wantRes: []string{},
		},
		{
			name:    "切片不为空",
			src:     []int{1, 2, 3},
			wantRes: []string{"1", "2", "3"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Map[int, string](tc.src, func(i int) string {
				return strconv.Itoa(i)
			})
			assert.Equal(t, res, tc.wantRes)
		})
	}
}

func TestMap_Filter(t *testing.T) {
	testCases := []struct {
		name    string
		src     []string
		wantRes []string
	}{
		{
			name:    "切片为空",
			src:     []string{},
			wantRes: []string{},
		},
		{
			name:    "切片不为空",
			src:     []string{"jack", "rocky", "mirror"},
			wantRes: []string{"rocky"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Filter[string](tc.src, func(e string) bool {
				return strings.HasPrefix(e, "r")
			})
			assert.Equal(t, res, tc.wantRes)
		})
	}
}
