package grokking

import (
	"reflect"
	"testing"
)

func TestTripletSumToZero(t *testing.T) {
	testcases := []struct {
		name           string
		inputArr       []int
		expectedOutput [][]int
	}{
		{
			name:           "TC1",
			inputArr:       []int{-3, 0, 1, 2, -1, 1, -2},
			expectedOutput: [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}},
		},
		{
			name:           "TC2",
			inputArr:       []int{-5, 2, -1, -2, 3},
			expectedOutput: [][]int{{-5, 2, 3}, {-2, -1, 3}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := tripletSumToZero(tc.inputArr)

			if !reflect.DeepEqual(actualResult, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actualResult)
			}
		})
	}
}
