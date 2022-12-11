package grokking

import (
	"reflect"
	"testing"
)

func TestTopoSort(t *testing.T) {
	r := topoSort(4, []Edge{
		{3, 2},
		{3, 0},
		{2, 0},
		{2, 1},
	})

	if !reflect.DeepEqual(r, []int{3, 2, 0, 1}) {
		panic("Wrong answer")
	}
}
