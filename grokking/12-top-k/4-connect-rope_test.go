package grokking

import (
	"reflect"
	"testing"
)

func Test_connectRopes(t *testing.T) {
	testcases := []struct {
		name           string
		ropesInput     []int
		expectedOutput int
	}{
		{"TC1", []int{1, 3, 11, 5}, 33},
		{"TC1", []int{3, 4, 5, 6}, 36},
		{"TC1", []int{1, 3, 11, 5, 2}, 42},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := connectRopes(tc.ropesInput)

			if !reflect.DeepEqual(actual, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
