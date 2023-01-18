package grokking

import "testing"

func TestFindKthSmallestFromMatrix(t *testing.T) {
	testcases := []struct {
		name           string
		inputMatrix    [][]int
		inputK         int
		expectedOutput int
	}{
		{
			name: "TC1",
			inputMatrix: [][]int{
				{2, 6, 7},
				{3, 7, 10},
				{5, 8, 11},
			},
			inputK:         5,
			expectedOutput: 7,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findKthSmallestFromMatrix(tc.inputMatrix, tc.inputK)

			if actual != tc.expectedOutput {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
