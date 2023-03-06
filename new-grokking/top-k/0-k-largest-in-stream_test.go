package topk

import (
	"reflect"
	"testing"
)

func Test_kthLargestInit(t *testing.T) {
	testcases := []struct {
		name           string
		inputFnCalls   []string
		inputArgs      [][]int
		outputExpected []int
	}{
		{
			name:           "TC1",
			inputFnCalls:   []string{"KthLargest", "Add", "Add", "Add", "Add", "Add"},
			inputArgs:      [][]int{{3, 4, 5, 8, 2}, {3}, {5}, {10}, {9}, {4}},
			outputExpected: []int{4, 5, 5, 8, 8},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var kthLargest KthLargest
			var actualOutput []int

			for i, call := range tc.inputFnCalls {
				if call == "KthLargest" {
					kthLargest.KthLargestInit(tc.inputArgs[i][0], tc.inputArgs[i][1:])
				} else if call == "Add" {
					kthLargest.Add(tc.inputArgs[i][0])

					result := kthLargest.ReturnKthLargest()

					actualOutput = append(actualOutput, result)
				} else {
					t.Error("Wrong fn call")
				}
			}

			if !reflect.DeepEqual(actualOutput, tc.outputExpected) {
				t.Errorf("Expected %v but got %v", tc.outputExpected, actualOutput)
			}
		})
	}
}
