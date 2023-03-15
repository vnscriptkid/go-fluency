package cyclicSort

import "testing"

func Test_findMissingNumber(t *testing.T) {
	testcases := []struct {
		Name           string
		InputArr       []int
		OutputExpected int
	}{
		{
			Name:           "TC1",
			InputArr:       []int{0, 1, 2, 4},
			OutputExpected: 3,
		},
		{
			Name:           "TC2",
			InputArr:       []int{3, 0, 1, 4},
			OutputExpected: 2,
		},
		{
			Name:           "TC3",
			InputArr:       []int{1, 4, 5, 6, 8, 2, 0, 7},
			OutputExpected: 3,
		},
		{
			Name:           "TC4",
			InputArr:       []int{1},
			OutputExpected: 0,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			outputActual := findMissingNumber(tc.InputArr)

			if outputActual != tc.OutputExpected {
				t.Errorf("Expected %v but got %v", tc.OutputExpected, outputActual)
			}
		})
	}
}
