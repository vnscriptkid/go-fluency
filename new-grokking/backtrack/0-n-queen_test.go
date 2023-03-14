package backtrack

import "testing"

func Test_nQueen(t *testing.T) {
	testcases := []struct {
		name           string
		inputN         int
		outputExpected int
	}{
		{
			name:           "TC1",
			inputN:         3,
			outputExpected: 0,
		},
		{
			name:           "TC2",
			inputN:         4,
			outputExpected: 2,
		},
		{
			name:           "TC3",
			inputN:         5,
			outputExpected: 10,
		},
		{
			name:           "TC4",
			inputN:         6,
			outputExpected: 4,
		},
		{
			name:           "TC5",
			inputN:         8,
			outputExpected: 92,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := nQueen(tc.inputN)

			if outputActual != tc.outputExpected {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
