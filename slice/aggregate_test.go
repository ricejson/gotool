package slice

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes int
		wantErr error
	}{
		{
			name:    "切片为nil",
			slice:   nil,
			wantRes: 0,
			wantErr: errors.New("slice is nil"),
		},
		{
			name:    "空切片",
			slice:   []int{},
			wantRes: 0,
			wantErr: errors.New("slice is empty"),
		},
		{
			name:    "正常情况",
			slice:   []int{5, 2, 1, 9, 2, 5, 4, 4},
			wantRes: 9,
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var slice = tc.slice
			res, err := Max[int](slice)
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes int
		wantErr error
	}{
		{
			name:    "切片为nil",
			slice:   nil,
			wantRes: 0,
			wantErr: errors.New("slice is nil"),
		},
		{
			name:    "空切片",
			slice:   []int{},
			wantRes: 0,
			wantErr: errors.New("slice is empty"),
		},
		{
			name:    "正常情况",
			slice:   []int{5, 2, 1, 9, 2, 5, 4, 4},
			wantRes: 1,
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var slice = tc.slice
			res, err := Min[int](slice)
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []complex128
		wantRes complex128
		wantErr error
	}{
		{
			name:    "切片为nil",
			slice:   nil,
			wantRes: complex(0, 0),
			wantErr: errors.New("slice is nil"),
		},
		{
			name:    "空切片",
			slice:   []complex128{},
			wantRes: complex(0, 0),
			wantErr: nil,
		},
		{
			name:    "正常情况",
			slice:   []complex128{complex(1, 2), complex(2, 3)},
			wantRes: complex(3, 5),
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var slice = tc.slice
			res, err := Sum[complex128](slice)
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
