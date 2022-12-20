package grokking

import "testing"

func TestEqualSubsetSumTopDown1(t *testing.T) {
	if r := equalSubsetSumTopDown([]int{1, 2, 3, 4}); r == false {
		panic("Expect true")
	}
}

func TestEqualSubsetSumTopDown2(t *testing.T) {
	if r := equalSubsetSumTopDown([]int{1, 1, 3, 4, 7}); r == false {
		panic("Expect true")
	}
}

func TestEqualSubsetSumTopDown3(t *testing.T) {
	if r := equalSubsetSumTopDown([]int{2, 3, 4, 6}); r == true {
		panic("Expect false")
	}
}
