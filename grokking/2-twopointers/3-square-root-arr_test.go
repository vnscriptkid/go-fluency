package grokking

import (
	"reflect"
	"testing"
)

func TestMakeSquares(t *testing.T) {
	testcases := []struct {
		inputArr []int
		expected []int
	}{
		{
			inputArr: []int{-2, -1, 0, 2, 3},
			expected: []int{0, 1, 4, 4, 9},
		},
		{
			inputArr: []int{-3, -1, 0, 1, 2},
			expected: []int{0, 1, 1, 4, 9},
		},
	}

	for _, tc := range testcases {
		actual := makeSquares(tc.inputArr)

		if !reflect.DeepEqual(tc.expected, actual) {
			t.Errorf("Expected %v but got %v", tc.expected, actual)
		}
	}
}
