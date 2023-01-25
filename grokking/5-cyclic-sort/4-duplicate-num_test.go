package grokking

import "testing"

func TestFindDuplicateNum(t *testing.T) {
	testcases := []struct {
		name           string
		inputArr       []int
		expectedOutput int
	}{
		{
			name:           "TC1",
			inputArr:       []int{1, 4, 4, 3, 2},
			expectedOutput: 4,
		},
		{
			name:           "TC2",
			inputArr:       []int{2, 1, 3, 3, 5, 4},
			expectedOutput: 3,
		},
		{
			name:           "TC1",
			inputArr:       []int{2, 4, 1, 4, 4},
			expectedOutput: 4,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findDuplicateNum(tc.inputArr)

			if actual != tc.expectedOutput {
				t.Errorf("Expected %d but got %d", tc.expectedOutput, actual)
			}
		})
	}
}
