package Track

import "testing"

func Test_permutePalindrome(t *testing.T) {
	testcases := []struct {
		name           string
		inputStr       string
		outputExpected bool
	}{
		{
			name:           "TC1",
			inputStr:       "peas",
			outputExpected: false,
		},
		{
			name:           "TC2",
			inputStr:       "abab",
			outputExpected: true,
		},
		{
			name:           "TC3",
			inputStr:       "racecar",
			outputExpected: true,
		},
		{
			name:           "TC4",
			inputStr:       "code",
			outputExpected: false,
		},
		{
			name:           "TC5",
			inputStr:       "baefeab",
			outputExpected: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := permutePalindrome(tc.inputStr)

			if actual != tc.outputExpected {
				t.Errorf("Expected %v but got %v", tc.outputExpected, actual)
			}
		})
	}
}
