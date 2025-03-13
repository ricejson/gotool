package pair

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPair_String(t *testing.T) {
	testCases := []struct {
		name       string
		key        any
		value      any
		wantString string
	}{
		{
			name:       "测试string",
			key:        "name",
			value:      "lisi",
			wantString: "<\"name\", \"lisi\">",
		},
		{
			name:       "测试slice",
			key:        "name",
			value:      []string{"a", "b", "c"},
			wantString: "<\"name\", []string{\"a\", \"b\", \"c\"}>",
		},
		{
			name: "测试map",
			key:  "name",
			value: map[string]int{
				"1": 1,
				"2": 2,
			},
			wantString: "<\"name\", map[string]int{\"1\":1, \"2\":2}>",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pair := NewPair[any, any](tc.key, tc.value)
			assert.Equal(t, pair.String(), tc.wantString)
		})
	}
}

func TestPair_Split(t *testing.T) {
	testCases := []struct {
		name      string
		key       any
		value     any
		wantKey   any
		wantValue any
	}{
		{
			name:      "测试string",
			key:       "name",
			value:     "lisi",
			wantKey:   "name",
			wantValue: "lisi",
		},
		{
			name:      "测试slice",
			key:       "name",
			value:     []string{"a", "b", "c"},
			wantKey:   "name",
			wantValue: []string{"a", "b", "c"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pair := NewPair[any, any](tc.key, tc.value)
			key, value := pair.Split()
			assert.Equal(t, key, tc.wantKey)
			assert.Equal(t, value, tc.wantValue)
		})
	}
}
