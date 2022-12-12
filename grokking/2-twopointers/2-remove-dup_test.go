package grokking

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	a := []int{2, 3, 3, 3, 6, 9, 9}

	removeDuplicates(a)

	fmt.Println(a)

	if !reflect.DeepEqual(a[0:4], []int{2, 3, 6, 9}) {
		panic("Wrong")
	}
}
