package grokking

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTopKLargestNumbers(t *testing.T) {
	r := topKLargestNumbers([]int{3, 1, 5, 12, 2, 11}, 3)

	if reflect.DeepEqual(r, []int{5, 11, 12}) == false {
		panic("Not the same arr")
	}

	fmt.Println(r)
}
