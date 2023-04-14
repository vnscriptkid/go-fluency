package kwaymerge

import (
	"reflect"
	"testing"
)

func Test_mergeSorted(t *testing.T) {
	testcases := []struct {
		name           string
		inputArr1      []int
		inputSize1     int
		inputArr2      []int
		inputSize2     int
		outputExpected []int
	}{
		{
			name:           "TC1",
			inputArr1:      []int{1, 2, 3, 0, 0, 0},
			inputSize1:     3,
			inputArr2:      []int{4, 5, 6},
			inputSize2:     3,
			outputExpected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:           "TC2",
			inputArr1:      []int{6, 7, 8, 9, 10, 0, 0, 0, 0, 0},
			inputSize1:     5,
			inputArr2:      []int{1, 2, 3, 4, 5},
			inputSize2:     5,
			outputExpected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := mergeSorted(tc.inputArr1, tc.inputSize1, tc.inputArr2, tc.inputSize2)

			if !reflect.DeepEqual(tc.outputExpected, outputActual) {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}

}
