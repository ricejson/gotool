package slice

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestIndexOf(t *testing.T) {
	testCases := []struct {
		name    string
		src     []string
		dst     string
		wantRes int
	}{
		{
			name:    "第一个元素符合条件",
			src:     []string{"a", "b", "a", "c"},
			dst:     "a",
			wantRes: 0,
		},
		{
			name:    "中间元素符合条件",
			src:     []string{"b", "b", "a", "c"},
			dst:     "a",
			wantRes: 2,
		},
		{
			name:    "没有元素符合条件",
			src:     []string{"b", "b", "b", "c"},
			dst:     "a",
			wantRes: -1,
		},
		{
			name:    "数组为空",
			src:     []string{},
			dst:     "a",
			wantRes: -1,
		},
		{
			name:    "数组为nil",
			src:     nil,
			dst:     "a",
			wantRes: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			assert.Equal(t, tc.wantRes, IndexOf[string](src, dst))
		})
	}
}

func TestIndexOfFunc(t *testing.T) {
	testCases := []struct {
		name      string
		src       []string
		dst       string
		compareTo func(e1, e2 string) bool
		wantRes   int
	}{
		{
			name: "第一个元素符合条件",
			src:  []string{"a", "b", "a", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: 0,
		},
		{
			name: "中间元素符合条件",
			src:  []string{"b", "b", "a", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: 2,
		},
		{
			name: "没有元素符合条件",
			src:  []string{"b", "b", "b", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: -1,
		},
		{
			name: "数组为空",
			src:  []string{},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: -1,
		},
		{
			name: "数组为nil",
			src:  nil,
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			var compareTo = tc.compareTo
			assert.Equal(t, tc.wantRes, IndexOfFunc[string](src, dst, compareTo))
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	testCases := []struct {
		name    string
		src     []string
		dst     string
		wantRes int
	}{
		{
			name:    "一个元素符合条件",
			src:     []string{"a", "b", "d", "c"},
			dst:     "a",
			wantRes: 0,
		},
		{
			name:    "多个元素符合条件",
			src:     []string{"a", "b", "a", "c"},
			dst:     "a",
			wantRes: 2,
		},
		{
			name:    "没有元素符合条件",
			src:     []string{"b", "b", "b", "c"},
			dst:     "a",
			wantRes: -1,
		},
		{
			name:    "数组为空",
			src:     []string{},
			dst:     "a",
			wantRes: -1,
		},
		{
			name:    "数组为nil",
			src:     nil,
			dst:     "a",
			wantRes: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			assert.Equal(t, tc.wantRes, LastIndexOf[string](src, dst))
		})
	}
}

func TestLastIndexOfFunc(t *testing.T) {
	testCases := []struct {
		name      string
		src       []string
		dst       string
		compareTo func(e1, e2 string) bool
		wantRes   int
	}{
		{
			name: "一个元素符合条件",
			src:  []string{"a", "b", "d", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: 0,
		},
		{
			name: "多个元素符合条件",
			src:  []string{"a", "b", "a", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: 2,
		},
		{
			name: "没有元素符合条件",
			src:  []string{"b", "b", "b", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: -1,
		},
		{
			name: "数组为空",
			src:  []string{},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: -1,
		},
		{
			name: "数组为nil",
			src:  nil,
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			var compareTo = tc.compareTo
			assert.Equal(t, tc.wantRes, LastIndexOfFunc[string](src, dst, compareTo))
		})
	}
}

func TestIndexOfAll(t *testing.T) {
	testCases := []struct {
		name    string
		src     []string
		dst     string
		wantRes []int
	}{
		{
			name:    "只有一个元素符合条件",
			src:     []string{"a", "b", "d", "c"},
			dst:     "a",
			wantRes: []int{0},
		},
		{
			name:    "多个元素符合条件",
			src:     []string{"a", "b", "a", "c"},
			dst:     "a",
			wantRes: []int{0, 2},
		},
		{
			name:    "没有元素符合条件",
			src:     []string{"b", "b", "b", "c"},
			dst:     "a",
			wantRes: []int{},
		},
		{
			name:    "数组为空",
			src:     []string{},
			dst:     "a",
			wantRes: []int{},
		},
		{
			name:    "数组为nil",
			src:     nil,
			dst:     "a",
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			assert.Equal(t, tc.wantRes, IndexOfAll[string](src, dst))
		})
	}
}

func TestIndexOfAllFunc(t *testing.T) {
	testCases := []struct {
		name      string
		src       []string
		dst       string
		compareTo func(e1, e2 string) bool
		wantRes   []int
	}{
		{
			name: "只有一个元素符合条件",
			src:  []string{"a", "b", "d", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: []int{0},
		},
		{
			name: "多个元素符合条件",
			src:  []string{"a", "b", "a", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: []int{0, 2},
		},
		{
			name: "没有元素符合条件",
			src:  []string{"b", "b", "b", "c"},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: []int{},
		},
		{
			name: "数组为空",
			src:  []string{},
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: []int{},
		},
		{
			name: "数组为nil",
			src:  nil,
			dst:  "a",
			compareTo: func(e1, e2 string) bool {
				return e1 == e2
			},
			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			var compareTo = tc.compareTo
			assert.Equal(t, tc.wantRes, IndexOfAllFunc[string](src, dst, compareTo))
		})
	}
}
