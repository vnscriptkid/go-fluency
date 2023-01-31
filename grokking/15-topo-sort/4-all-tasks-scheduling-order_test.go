package grokking

import (
	"reflect"
	"testing"
)

func Test_allTasksSchedulingOrder(t *testing.T) {
	testcases := []struct {
		name               string
		inputTasks         int
		inputPrerequisites [][2]int
		expectedOutput     [][]int
	}{
		// {
		// 	name:               "TC1",
		// 	inputTasks:         3,
		// 	inputPrerequisites: [][2]int{{0, 1}, {1, 2}},
		// 	expectedOutput:     [][]int{{0, 1, 2}},
		// },
		{
			name:               "TC2",
			inputTasks:         4,
			inputPrerequisites: [][2]int{{3, 2}, {3, 0}, {2, 0}, {2, 1}},
			expectedOutput:     [][]int{{3, 2, 0, 1}, {3, 2, 1, 0}},
		},
		{
			name:               "TC3",
			inputTasks:         6,
			inputPrerequisites: [][2]int{{2, 5}, {0, 5}, {0, 4}, {1, 4}, {3, 2}, {1, 3}},
			expectedOutput: [][]int{
				{0, 1, 4, 3, 2, 5},
				{0, 1, 3, 4, 2, 5},
				{0, 1, 3, 2, 4, 5},
				{0, 1, 3, 2, 5, 4},
				{1, 0, 3, 4, 2, 5},
				{1, 0, 3, 2, 4, 5},
				{1, 0, 3, 2, 5, 4},
				{1, 0, 4, 3, 2, 5},
				{1, 3, 0, 2, 4, 5},
				{1, 3, 0, 2, 5, 4},
				{1, 3, 0, 4, 2, 5},
				{1, 3, 2, 0, 5, 4},
				{1, 3, 2, 0, 4, 5},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := allTasksSchedulingOrder(tc.inputTasks, tc.inputPrerequisites)

			if !reflect.DeepEqual(actual, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
