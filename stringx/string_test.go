package stringx

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestString_IsEmpty(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantRes bool
	}{
		{
			name:    "string为空字符串",
			input:   "",
			wantRes: true,
		},
		{
			name:    "string不为空但都是空字符",
			input:   "\t\r\n",
			wantRes: false,
		},
		{
			name:    "string为空白字符",
			input:   "   ",
			wantRes: false,
		},
		{
			name:    "string不为空串也不为空字符",
			input:   "  abc",
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantRes, IsEmpty(tc.input))
		})
	}
}

func TestString_IsNotEmpty(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantRes bool
	}{
		{
			name:    "string为空字符串",
			input:   "",
			wantRes: false,
		},
		{
			name:    "string不为空但都是空字符",
			input:   "\t\r\n",
			wantRes: true,
		},
		{
			name:    "string为空白字符",
			input:   "   ",
			wantRes: true,
		},
		{
			name:    "string不为空串也不为空字符",
			input:   "  abc",
			wantRes: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantRes, IsNotEmpty(tc.input))
		})
	}
}

func TestString_IsBlank(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantRes bool
	}{
		{
			name:    "string为空字符串",
			input:   "",
			wantRes: true,
		},
		{
			name:    "string不为空但都是空字符",
			input:   "\t\r\n",
			wantRes: true,
		},
		{
			name:    "string为空白字符",
			input:   "   ",
			wantRes: true,
		},
		{
			name:    "string不为空串也不为空字符",
			input:   "  abc",
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantRes, IsBlank(tc.input))
		})
	}
}

func TestString_IsNotBlank(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantRes bool
	}{
		{
			name:    "string为空字符串",
			input:   "",
			wantRes: false,
		},
		{
			name:    "string不为空但都是空字符",
			input:   "\t\r\n",
			wantRes: false,
		},
		{
			name:    "string为空白字符",
			input:   "   ",
			wantRes: false,
		},
		{
			name:    "string不为空串也不为空字符",
			input:   "  abc",
			wantRes: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantRes, IsNotBlank(tc.input))
		})
	}
}
