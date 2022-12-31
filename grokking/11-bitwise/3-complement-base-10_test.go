package grokking

import "testing"

func TestBitwiseComplement(t *testing.T) {
	testcases := []struct {
		name           string
		input          int
		expectedOutput int
	}{
		{
			name:           "TC1",
			input:          8,
			expectedOutput: 7,
		},
		{
			name:           "TC2",
			input:          10,
			expectedOutput: 5,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := bitwiseComplement(tc.input)

			if actualOutput != tc.expectedOutput {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actualOutput)
			}
		})
	}
}
