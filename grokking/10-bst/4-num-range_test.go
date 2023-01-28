package grokking

import (
	"reflect"
	"testing"
)

func Test_numberRange(t *testing.T) {
	testcases := []struct {
		name           string
		inputArr       []int
		inputKey       int
		expectedOutput [2]int
	}{
		{
			name:           "TC1",
			inputArr:       []int{4, 6, 6, 6, 9},
			inputKey:       6,
			expectedOutput: [2]int{1, 3},
		},
		{
			name:           "TC2",
			inputArr:       []int{1, 3, 8, 10, 15},
			inputKey:       10,
			expectedOutput: [2]int{3, 3},
		},
		{
			name:           "TC3",
			inputArr:       []int{1, 3, 8, 10, 15},
			inputKey:       12,
			expectedOutput: [2]int{-1, -1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := numberRange(tc.inputArr, tc.inputKey)

			if !reflect.DeepEqual(actual, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
