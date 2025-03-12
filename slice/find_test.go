package slice

import (
	"github.com/go-playground/assert/v2"
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		name       string
		src        []string
		matchFunc  func(string) bool
		wantRes    string
		wantExists bool
	}{
		{
			name: "src中没有符合条件的",
			src:  []string{"aaa", "bbb", "ccc"},
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes:    "",
			wantExists: false,
		},
		{
			name: "src中有一个符合条件的",
			src:  []string{"aaa", "bbb", "ccc", "ddd"},
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes:    "ddd",
			wantExists: true,
		},
		{
			name: "src中有多个符合条件的",
			src:  []string{"aaa", "bbbdd", "ccc", "ddd"},
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes:    "bbbdd",
			wantExists: true,
		},
		{
			name: "src为空",
			src:  []string{},
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes:    "",
			wantExists: false,
		},
		{
			name: "src为nil",
			src:  nil,
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes:    "",
			wantExists: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			res, exists := Find[string](src, tc.matchFunc)
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantExists, exists)
		})
	}
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name      string
		src       []string
		matchFunc func(string) bool
		wantRes   []string
	}{
		{
			name: "src中没有符合条件的",
			src:  []string{"aaa", "bbb", "ccc"},
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes: []string{},
		},
		{
			name: "src中有一个符合条件的",
			src:  []string{"aaa", "bbb", "ccc", "ddd"},
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes: []string{"ddd"},
		},
		{
			name: "src中有多个符合条件的",
			src:  []string{"aaa", "bbbdd", "ccc", "ddd"},
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes: []string{"bbbdd", "ddd"},
		},
		{
			name: "src为空",
			src:  []string{},
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes: []string{},
		},
		{
			name: "src为nil",
			src:  nil,
			matchFunc: func(s string) bool {
				return strings.Contains(s, "dd")
			},
			wantRes: []string{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var src = tc.src
			res := FindAll[string](src, tc.matchFunc)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
