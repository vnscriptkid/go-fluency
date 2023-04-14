package stack

import "testing"

func Test_calculator(t *testing.T) {
	testcases := []struct {
		name           string
		input          string
		outputExpected int
	}{
		{
			name:           "TC1",
			input:          "12 - (6 + 2) + 5",
			outputExpected: 9,
		},
		{
			name:           "TC2",
			input:          "(8 + 100) + (13 - 8 - (2 + 1))",
			outputExpected: 110,
		},
		{
			name:           "TC3",
			input:          "40 - 25 - 5",
			outputExpected: 10,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := calculator(tc.input)

			if outputActual != tc.outputExpected {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
