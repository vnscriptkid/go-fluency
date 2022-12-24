package grokking

import "testing"

func TestFindHappyNum(t *testing.T) {
	testcases := []struct {
		name     string
		inputNum int
		expected bool
	}{
		{
			name:     "TC1",
			inputNum: 23,
			expected: true,
		},
		{
			name:     "TC2",
			inputNum: 12,
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findHappyNum(tc.inputNum)

			if actual != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
