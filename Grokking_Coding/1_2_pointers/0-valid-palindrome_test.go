package TwoPointers

import "testing"

func Test_isPalindrome(t *testing.T) {
	testcases := []struct {
		Name           string
		InputString    string
		OutputExpected bool
	}{
		{
			Name:           "Test Case 1",
			InputString:    "abcdcba",
			OutputExpected: true,
		},
		{
			Name:           "Test Case 2",
			InputString:    "a",
			OutputExpected: true,
		},
		{
			Name:           "Test Case 3",
			InputString:    "ab",
			OutputExpected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			outputActual := isPalindrome(tc.InputString)

			if outputActual != tc.OutputExpected {
				t.Errorf("Expected %v, got %v", tc.OutputExpected, outputActual)
			}
		})
	}
}
