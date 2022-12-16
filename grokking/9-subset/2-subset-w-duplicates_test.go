package grokking

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindSubsetsWDuplicates(t *testing.T) {
	r := findSubsetsWDuplicates([]int{1, 3, 3})

	fmt.Println(r)

	if !reflect.DeepEqual(r, [][]int{{}, {1}, {3}, {1, 3}, {3, 3}, {1, 3, 3}}) {
		panic("Hix")
	}
}
