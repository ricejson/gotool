package gotool

import (
	"fmt"
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

func TestCondition_IfThenElseFunc(t *testing.T) {
	testCases := []struct {
		name      string
		condition bool
		trueFunc  func() (string, error)
		falseFunc func() (string, error)
		wantRes   string
		wantErr   error
	}{
		{
			name:      "条件成立",
			condition: true,
			trueFunc: func() (string, error) {
				fmt.Println("条件成立")
				return "true", nil
			},
			falseFunc: func() (string, error) {
				fmt.Println("条件不成立")
				return "false", nil
			},
			wantRes: "true",
			wantErr: nil,
		},
		{
			name:      "条件不成立",
			condition: false,
			trueFunc: func() (string, error) {
				fmt.Println("条件成立")
				return "true", nil
			},
			falseFunc: func() (string, error) {
				fmt.Println("条件不成立")
				return "false", nil
			},
			wantRes: "false",
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IfThenElseFunc[string](tc.condition, tc.trueFunc, tc.falseFunc)
			assert.Equal(t, res, tc.wantRes)
			assert.Equal(t, err, tc.wantErr)
		})
	}
}
