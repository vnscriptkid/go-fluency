package grokking

import "testing"

func TestNextLetter(t *testing.T) {
	testcases := []struct {
		name           string
		inputLetters   []rune
		inputKey       rune
		expectedOutput rune
	}{
		{
			name:           "TC1",
			inputLetters:   []rune{'a', 'c', 'f', 'h'},
			inputKey:       'f',
			expectedOutput: 'h',
		},
		{
			name:           "TC2",
			inputLetters:   []rune{'a', 'c', 'f', 'h'},
			inputKey:       'b',
			expectedOutput: 'c',
		},
		{
			name:           "TC3",
			inputLetters:   []rune{'a', 'c', 'f', 'h'},
			inputKey:       'm',
			expectedOutput: 'a',
		},
		{
			name:           "TC4",
			inputLetters:   []rune{'a', 'c', 'f', 'h'},
			inputKey:       'h',
			expectedOutput: 'a',
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := nextLetter(tc.inputLetters, tc.inputKey)

			if actualOutput != tc.expectedOutput {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actualOutput)
			}
		})
	}
}
