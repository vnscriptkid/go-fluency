package greedy

import (
	"reflect"
	"testing"
)

func Test_jumpGame(t *testing.T) {
	testcases := []struct {
		name           string
		inputArr       []int
		outputExpected bool
	}{
		{
			name:           "TC1",
			inputArr:       []int{2, 3, 1, 1, 9},
			outputExpected: true,
		},
		{
			name:           "TC2",
			inputArr:       []int{3, 2, 1, 0, 4},
			outputExpected: false,
		},
		{
			name:           "TC3",
			inputArr:       []int{4, 0, 0, 0, 4},
			outputExpected: true,
		},
		{
			name:           "TC4",
			inputArr:       []int{0},
			outputExpected: true,
		},
		{
			name:           "TC5",
			inputArr:       []int{1},
			outputExpected: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := jumpGame(tc.inputArr)

			if !reflect.DeepEqual(outputActual, tc.outputExpected) {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
