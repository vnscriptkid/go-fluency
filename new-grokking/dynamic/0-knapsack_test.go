package dynamic

import "testing"

func Test_findMaxKnapsackProfit(t *testing.T) {
	testcases := []struct {
		Name           string
		InputCap       int
		InputWeights   []int
		InputValues    []int
		OutputExpected int
	}{
		{
			Name:           "TC1",
			InputCap:       6,
			InputWeights:   []int{1, 2, 3, 5},
			InputValues:    []int{1, 5, 4, 8},
			OutputExpected: 10,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			outputActual := findMaxKnapsackProfit(tc.InputCap, tc.InputWeights, tc.InputValues)

			if outputActual != tc.OutputExpected {
				t.Errorf("Expected %v but got %v", tc.OutputExpected, outputActual)
			}
		})
	}
}
