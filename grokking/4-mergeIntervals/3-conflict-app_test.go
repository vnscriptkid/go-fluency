package grokking

import "testing"

func TestCanAttendAllApp(t *testing.T) {
	testcases := []struct {
		name     string
		inputArg [][2]int
		expected bool
	}{
		{
			name:     "TC1",
			inputArg: [][2]int{{1, 4}, {2, 5}, {7, 9}},
			expected: false,
		},
		{
			name:     "TC2",
			inputArg: [][2]int{{6, 7}, {2, 4}, {8, 12}},
			expected: true,
		},
		{
			name:     "TC3",
			inputArg: [][2]int{{4, 5}, {2, 3}, {3, 6}},
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := canAttendAllApp(tc.inputArg)

			if actual != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
