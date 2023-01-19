package grokking

import "testing"

func TestLongestSubstrKDistinctChars(t *testing.T) {
	testcases := []struct {
		name           string
		inputStr       string
		inputK         int
		expectedOutput int
	}{
		{
			name:           "TC1",
			inputStr:       "araaci",
			inputK:         2,
			expectedOutput: 4,
		},
		{
			name:           "TC2",
			inputStr:       "araaci",
			inputK:         1,
			expectedOutput: 2,
		},
		{
			name:           "TC3",
			inputStr:       "cbbebi",
			inputK:         3,
			expectedOutput: 5,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := longestSubstrKDistinctChars(tc.inputStr, tc.inputK)

			if actual != tc.expectedOutput {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
