package grokking

import (
	"reflect"
	"testing"
)

func Test_flipThenInvert(t *testing.T) {
	testcases := []struct {
		name           string
		inputMatrix    [][]int
		expectedOutput [][]int
	}{
		{
			name: "TC1",
			inputMatrix: [][]int{
				{1, 0, 1},
				{1, 1, 1},
				{0, 1, 1},
			},
			expectedOutput: [][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
		},
		{
			name: "TC2",
			inputMatrix: [][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
			expectedOutput: [][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
				{1, 0, 1, 0},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := flipThenInvert(tc.inputMatrix)

			if !reflect.DeepEqual(actual, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
