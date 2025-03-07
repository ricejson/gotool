package slice

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  int
		want bool
	}{
		{
			name: "dst不存在",
			src:  []int{1, 3, 5},
			dst:  4,
			want: false,
		},
		{
			name: "dst存在",
			src:  []int{1, 3, 5},
			dst:  5,
			want: true,
		},
		{
			name: "src长度为0",
			src:  []int{},
			dst:  2,
			want: false,
		},
		{
			name: "src为nil",
			src:  nil,
			dst:  2,
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			res := Contains[int](src, dst)
			assert.Equal(t, res, tc.want)
		})
	}
}

func TestContainsFunc(t *testing.T) {
	testCases := []struct {
		name      string
		src       []int
		dst       int
		compareTo func(e1, e2 int) bool
		want      bool
	}{
		{
			name: "dst不存在",
			src:  []int{1, 3, 5},
			dst:  4,
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: false,
		},
		{
			name: "dst存在",
			src:  []int{1, 3, 5},
			dst:  5,
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: true,
		},
		{
			name: "src为空",
			src:  []int{},
			dst:  4,
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: false,
		},
		{
			name: "src为nil",
			src:  nil,
			dst:  4,
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			var compareTo = tc.compareTo
			var res = ContainsFunc[int](src, dst, compareTo)
			assert.Equal(t, res, tc.want)
		})
	}
}

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want bool
	}{
		{
			name: "src存在dst中某些元素",
			src:  []int{1, 3, 5},
			dst:  []int{1, 4, 7},
			want: true,
		},
		{
			name: "src存在dst中全部元素",
			src:  []int{1, 3, 5},
			dst:  []int{1, 3, 5},
			want: true,
		},
		{
			name: "src不存在dst中全部元素",
			src:  []int{1, 3, 5},
			dst:  []int{2, 4, 7},
			want: false,
		},
		{
			name: "src为nil",
			src:  nil,
			dst:  []int{1, 4, 7},
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			var res = ContainsAny[int](src, dst)
			assert.Equal(t, res, tc.want)
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	testCases := []struct {
		name      string
		src       []int
		dst       []int
		compareTo func(e1, e2 int) bool
		want      bool
	}{
		{
			name: "src存在dst中某些元素",
			src:  []int{1, 3, 5},
			dst:  []int{1, 4, 7},
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: true,
		},
		{
			name: "src存在dst中全部元素",
			src:  []int{1, 3, 5},
			dst:  []int{1, 3, 5},
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: true,
		},
		{
			name: "src不存在dst中全部元素",
			src:  []int{1, 3, 5},
			dst:  []int{2, 4, 7},
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: false,
		},
		{
			name: "src为nil",
			src:  nil,
			dst:  []int{1, 4, 7},
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			var compareTo = tc.compareTo
			var res = ContainsAnyFunc[int](src, dst, compareTo)
			assert.Equal(t, res, tc.want)
		})
	}
}

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want bool
	}{
		{
			name: "src存在dst中某些元素",
			src:  []int{1, 3, 5, 8},
			dst:  []int{1, 4, 7},
			want: false,
		},
		{
			name: "src存在dst中全部元素",
			src:  []int{1, 3, 5, 8},
			dst:  []int{1, 3, 5},
			want: true,
		},
		{
			name: "src不存在dst中全部元素",
			src:  []int{1, 3, 5},
			dst:  []int{2, 4, 7},
			want: false,
		},
		{
			name: "src为nil",
			src:  nil,
			dst:  []int{1, 4, 7},
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			var res = ContainsAll[int](src, dst)
			assert.Equal(t, res, tc.want)
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	testCases := []struct {
		name      string
		src       []int
		dst       []int
		compareTo func(e1, e2 int) bool
		want      bool
	}{
		{
			name: "src存在dst中某些元素",
			src:  []int{1, 3, 5, 8},
			dst:  []int{1, 4, 7},
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: false,
		},
		{
			name: "src存在dst中全部元素",
			src:  []int{1, 3, 5, 8},
			dst:  []int{1, 3, 5},
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: true,
		},
		{
			name: "src不存在dst中全部元素",
			src:  []int{1, 3, 5},
			dst:  []int{2, 4, 7},
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: false,
		},
		{
			name: "src为nil",
			src:  nil,
			dst:  []int{1, 4, 7},
			compareTo: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			var dst = tc.dst
			var compareTo = tc.compareTo
			var res = ContainsAllFunc[int](src, dst, compareTo)
			assert.Equal(t, res, tc.want)
		})
	}
}
