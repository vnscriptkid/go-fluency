package grokking

import (
	"reflect"
	"testing"
)

func Test_smallestNumRange(t *testing.T) {
	testcases := []struct {
		name            string
		sortedArrsInput [][]int
		expectedOutput  [2]int
	}{
		{
			"TC1",
			[][]int{{1, 5, 8}, {4, 12}, {7, 8, 10}},
			[2]int{4, 7},
		},
		{
			"TC2",
			[][]int{{1, 9}, {4, 12}, {7, 10, 16}},
			[2]int{9, 12},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := smallestNumRange(tc.sortedArrsInput)

			if !reflect.DeepEqual(actual, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
