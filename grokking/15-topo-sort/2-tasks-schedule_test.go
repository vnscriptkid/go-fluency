package grokking

import "testing"

func TestTasksSchedule1(t *testing.T) {
	r := taskScheduling(3, [][2]int{{0, 1}, {1, 2}})

	if r != true {
		panic("Expect true")
	}
}

func TestTasksSchedule2(t *testing.T) {
	r := taskScheduling(3, [][2]int{{0, 1}, {1, 2}, {2, 0}})

	if r != false {
		panic("Expect false")
	}
}

func TestTasksSchedule3(t *testing.T) {
	r := taskScheduling(6, [][2]int{{2, 5}, {0, 5}, {0, 4}, {1, 4}, {3, 2}, {1, 3}})

	if r != true {
		panic("Expect true")
	}
}
