package grokking

import (
	"reflect"
	"testing"
)

func Test_permutationChangingCase(t *testing.T) {
	testcases := []struct {
		name           string
		inputStr       string
		expectedOutput []string
	}{
		{
			name:           "TC1",
			inputStr:       "ad52",
			expectedOutput: []string{"ad52", "Ad52", "aD52", "AD52"},
		},
		{
			name:           "TC2",
			inputStr:       "ab7c",
			expectedOutput: []string{"ab7c", "Ab7c", "aB7c", "AB7c", "ab7C", "Ab7C", "aB7C", "AB7C"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := permutationChangingCase(tc.inputStr)

			if !reflect.DeepEqual(actual, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
