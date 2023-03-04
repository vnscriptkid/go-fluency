package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_findMaxSlidingWindow(t *testing.T) {
	testcases := []struct {
		name            string
		inputArr        []int
		inputWindowSize int
		outputExpected  []int
	}{
		{
			name:            "TC1",
			inputArr:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputWindowSize: 3,
			outputExpected:  []int{3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := findMaxSlidingWindow(tc.inputArr, tc.inputWindowSize)

			if !reflect.DeepEqual(outputActual, tc.outputExpected) {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
