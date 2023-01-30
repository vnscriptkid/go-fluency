package grokking

import (
	"testing"
)

func Test_minSubsetSumDiff(t *testing.T) {
	testcases := []struct {
		name         string
		inputArr     []int
		expectOutput int
	}{
		{"TC1", []int{1, 2, 3, 9}, 3},
		{"TC2", []int{1, 2, 7, 1, 5}, 0},
		{"TC3", []int{1, 3, 100, 4}, 92},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := minSubsetSumDiff(tc.inputArr)

			if actual != tc.expectOutput {
				t.Errorf("Expected %v but got %v", tc.expectOutput, actual)
			}
		})
	}
}
