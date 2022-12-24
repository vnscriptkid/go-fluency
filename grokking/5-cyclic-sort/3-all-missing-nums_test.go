package grokking

import (
	"reflect"
	"testing"
)

func TestFindAllMissingNumbers(t *testing.T) {
	testcases := []struct {
		name       string
		inputSlice []int
		expected   []int
	}{
		{
			name:       "TC1",
			inputSlice: []int{2, 3, 1, 8, 2, 3, 5, 1},
			expected:   []int{4, 6, 7},
		},
		{
			name:       "TC2",
			inputSlice: []int{2, 4, 1, 2},
			expected:   []int{3},
		},
		{
			name:       "TC3",
			inputSlice: []int{2, 3, 2, 1},
			expected:   []int{4},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findAllMissingNums(tc.inputSlice)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
