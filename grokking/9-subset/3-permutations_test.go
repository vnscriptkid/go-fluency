package grokking

import (
	"reflect"
	"testing"
)

func TestFindPermutations(t *testing.T) {
	r := findPermutations([]int{1, 3, 5})

	expected := [][]int{
		{5, 3, 1},
		{3, 5, 1},
		{3, 1, 5},
		{5, 1, 3},
		{1, 5, 3},
		{1, 3, 5},
	}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected %v but got %v", expected, r)
	}
}
