package gotool

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestPtr_ToPtr(t *testing.T) {
	testCases := []struct {
		name    string
		src     any
		wantRes any
	}{
		{
			name:    "测试string类型",
			src:     "Hello World",
			wantRes: "Hello World",
		},
		{
			name:    "测试bool类型",
			src:     true,
			wantRes: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, *ToPtr[any](tc.src), tc.wantRes)
		})
	}
}
