package subsets

import (
	"reflect"
	"testing"
)

// findAllSubsets will find all the subsets of the given array
func Test_findAllSubsets(t *testing.T) {
	testcases := []struct {
		name           string
		inputArr       []int
		outputExpected [][]int
	}{
		{
			name:           "TC1",
			inputArr:       []int{1},
			outputExpected: [][]int{{}, {1}},
		},
		{
			name:     "TC2",
			inputArr: []int{2, 5, 7},
			outputExpected: [][]int{
				{},
				{2},
				{5},
				{2, 5},
				{7},
				{2, 7},
				{5, 7},
				{2, 5, 7},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := findAllSubsets(tc.inputArr)

			if !reflect.DeepEqual(outputActual, tc.outputExpected) {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
