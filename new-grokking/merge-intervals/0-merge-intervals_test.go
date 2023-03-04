package mergeintervals

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	testcases := []struct {
		name           string
		inputIntervals []Interval
		outputExpected []Interval
	}{
		{
			name: "TC1",
			inputIntervals: []Interval{
				{start: 1, end: 5},
				{start: 3, end: 7},
				{start: 4, end: 6},
			},
			outputExpected: []Interval{
				{start: 1, end: 7},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := mergeIntervals(tc.inputIntervals)

			if !reflect.DeepEqual(outputActual, tc.outputExpected) {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
