package grokking

import "testing"

func TestCeilingOfNum1(t *testing.T) {
	r := ceilingOfNum([]int{4, 6, 10}, 6)

	if r != 1 {
		panic("Expect 1")
	}
}

func TestCeilingOfNum2(t *testing.T) {
	r := ceilingOfNum([]int{1, 3, 8, 10, 15}, 12)

	if r != 4 {
		panic("Expect 4")
	}
}

func TestCeilingOfNum3(t *testing.T) {
	r := ceilingOfNum([]int{4, 6, 10}, 17)

	if r != -1 {
		panic("Expect -1")
	}
}

func TestCeilingOfNum4(t *testing.T) {
	r := ceilingOfNum([]int{4, 6, 10}, 2)

	if r != 0 {
		panic("Expect 0")
	}
}
