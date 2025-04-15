package gotool

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCondition_IfThenElse(t *testing.T) {
	testCases := []struct {
		name       string
		condition  bool
		trueValue  string
		falseValue string
		wantRes    string
	}{
		{
			name:       "条件成立",
			condition:  true,
			trueValue:  "true",
			falseValue: "false",
			wantRes:    "true",
		},
		{
			name:       "条件不成立",
			condition:  false,
			trueValue:  "true",
			falseValue: "false",
			wantRes:    "false",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IfThenElse[string](tc.condition, tc.trueValue, tc.falseValue)
			assert.Equal(t, res, tc.wantRes)
		})
	}
}
