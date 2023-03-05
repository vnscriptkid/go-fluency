package twoheaps

import "testing"

func Test_maximumCapital(t *testing.T) {
	testcases := []struct {
		name           string
		inputCapital   int
		inputK         int
		inputCapitals  []int
		inputProfits   []int
		outputExpected int
	}{
		{
			name:           "TC1",
			inputCapital:   1,
			inputK:         2,
			inputCapitals:  []int{1, 2, 2, 3},
			inputProfits:   []int{2, 4, 6, 8},
			outputExpected: 11,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := maximumCapital(tc.inputCapital, tc.inputK, tc.inputCapitals, tc.inputProfits)

			if outputActual != tc.outputExpected {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
