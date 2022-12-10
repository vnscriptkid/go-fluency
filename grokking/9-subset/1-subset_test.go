package grokking

import (
	"fmt"
	"testing"
)

func TestSubset(t *testing.T) {
	r := subset([]int{1, 3})

	fmt.Println(r)

	r = subset([]int{1, 3, 5})

	fmt.Println(r)
}
