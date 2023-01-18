package grokking

import "testing"

func TestSubsetSumTopdown(t *testing.T) {
	testcases := []struct {
		name           string
		inputSubset    []int
		inputS         int
		expectedOutput bool
	}{
		{
			name:           "TC1",
			inputSubset:    []int{1, 2, 3, 7},
			inputS:         6,
			expectedOutput: true,
		},
		{
			name:           "TC2",
			inputSubset:    []int{1, 2, 7, 1, 5},
			inputS:         10,
			expectedOutput: true,
		},
		{
			name:           "TC3",
			inputSubset:    []int{1, 3, 4, 8},
			inputS:         6,
			expectedOutput: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := subsetSumTopdown(tc.inputSubset, tc.inputS)

			if actual != tc.expectedOutput {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
