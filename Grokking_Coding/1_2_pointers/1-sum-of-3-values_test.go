package TwoPointers

import "testing"

func Test_findSumOfThree(t *testing.T) {
	testcases := []struct {
		Name           string
		InputNums      []int
		InputTarget    int
		OutputExpected bool
	}{
		{
			Name:           "Test Case 1",
			InputNums:      []int{1, 2, 3, 4, 5},
			InputTarget:    9,
			OutputExpected: true,
		},
		{
			Name:           "Test Case 2",
			InputNums:      []int{1, 2, 3, 4, 5},
			InputTarget:    2,
			OutputExpected: false,
		},
		{
			Name:           "Test Case 3",
			InputNums:      []int{1, 2, 3, 4, 5},
			InputTarget:    15,
			OutputExpected: false,
		},
		{
			Name:           "Test Case 4",
			InputNums:      []int{1, 2, 3, 4, 5},
			InputTarget:    8,
			OutputExpected: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			outputActual := findSumOfThree(tc.InputNums, tc.InputTarget)

			if outputActual != tc.OutputExpected {
				t.Errorf("Expected %v, got %v", tc.OutputExpected, outputActual)
			}
		})
	}
}
