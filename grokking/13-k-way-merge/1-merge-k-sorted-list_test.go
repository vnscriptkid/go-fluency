package grokking

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeKSortedLists(t *testing.T) {
	r := mergeKSortedLists([][]int{
		{2, 6, 8},
		{3, 6, 7},
		{1, 3, 4},
	})

	if !reflect.DeepEqual(r, []int{1, 2, 3, 3, 4, 6, 6, 7, 8}) {
		panic("Wrong result")
	}

	fmt.Println("Correct result: ", r)
}
