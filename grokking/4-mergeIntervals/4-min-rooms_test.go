package grokking

import (
	"testing"
)

func TestMinMeetingRooms(t *testing.T) {
	testcases := []struct {
		name           string
		meetingsInput  [][2]int
		expectedOutput int
	}{
		{
			name:           "TC1",
			meetingsInput:  [][2]int{{1, 4}, {2, 5}, {7, 9}},
			expectedOutput: 2,
		},
		{
			name:           "TC2",
			meetingsInput:  [][2]int{{6, 7}, {2, 4}, {8, 12}},
			expectedOutput: 1,
		},
		{
			name:           "TC3",
			meetingsInput:  [][2]int{{1, 4}, {2, 3}, {3, 6}},
			expectedOutput: 2,
		},
		{
			name:           "TC4",
			meetingsInput:  [][2]int{{4, 5}, {2, 3}, {2, 4}, {3, 5}},
			expectedOutput: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := minMeetingRooms(tc.meetingsInput)

			if actual != tc.expectedOutput {
				t.Errorf("Expected %d but got %d", tc.expectedOutput, actual)
			}
		})
	}
}
