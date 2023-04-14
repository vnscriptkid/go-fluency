package bst

import "testing"

func Test_binarySearchRotated(t *testing.T) {
	testcases := []struct {
		name           string
		inputArr       []int
		inputTarget    int
		outputExpected int
	}{
		{
			name:           "TC1",
			inputArr:       []int{6, 7, 1, 2, 3, 4, 5},
			inputTarget:    3,
			outputExpected: 4,
		},
		{
			name:           "TC2",
			inputArr:       []int{6, 7, 1, 2, 3, 4, 5},
			inputTarget:    6,
			outputExpected: 0,
		},
		{
			name:           "TC3",
			inputArr:       []int{4, 5, 6, 1, 2, 3},
			inputTarget:    3,
			outputExpected: 5,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := binarySearchRotated(tc.inputArr, tc.inputTarget)

			if outputActual != tc.outputExpected {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
