package grokking

import (
	"reflect"
	"testing"
)

func TestTasksSchedulingOrder(t *testing.T) {
	testcases := []struct {
		name               string
		tasksInput         int
		prerequisitesInput [][2]int
		expectedOutput     []int
	}{
		{
			name:               "TC1",
			tasksInput:         3,
			prerequisitesInput: [][2]int{{0, 1}, {1, 2}},
			expectedOutput:     []int{0, 1, 2},
		},
		{
			name:               "TC2",
			tasksInput:         3,
			prerequisitesInput: [][2]int{{0, 1}, {1, 2}, {2, 0}},
			expectedOutput:     []int{},
		},
		{
			name:               "TC3",
			tasksInput:         6,
			prerequisitesInput: [][2]int{{2, 5}, {0, 5}, {0, 4}, {1, 4}, {3, 2}, {1, 3}},
			expectedOutput:     []int{0, 1, 4, 3, 2, 5},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tasksSchedulingOrder(tc.tasksInput, tc.prerequisitesInput)

			if !reflect.DeepEqual(actual, tc.expectedOutput) {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
