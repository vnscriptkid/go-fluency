package grokking

import (
	"reflect"
	"testing"
)

func TestFindKClosestPoints(t *testing.T) {
	testcases := []struct {
		name           string
		inputPoints    [][2]int
		inputK         int
		expectedOutput [][2]int
	}{
		{
			name:           "TC1",
			inputPoints:    [][2]int{{1, 2}, {1, 3}},
			inputK:         1,
			expectedOutput: [][2]int{{1, 2}},
		},
		{
			name:           "TC2",
			inputPoints:    [][2]int{{1, 3}, {3, 4}, {2, -1}},
			inputK:         2,
			expectedOutput: [][2]int{{1, 3}, {2, -1}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := findKClosestPoints(tc.inputPoints, tc.inputK)

			if !reflect.DeepEqual(actualOutput, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actualOutput)
			}

		})
	}
}
