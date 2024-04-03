package main

import "fmt"

func main() {
	// Slice of string set to its zero value state.
	// zero value for a slice is nil.
	var slice1 []string
	// Slice of string set to its empty state.
	slice2 := []string{}
	// Slice of string set with a length and capacity of 5.
	slice3 := make([]string, 5)
	// Slice of string set with a length of 5 and capacity of 8.
	slice4 := make([]string, 5, 8)
	// Slice of string set with values with a length and capacity of 5.
	slice5 := []string{"A", "B", "C", "D", "E"}

	// %#v: a Go-syntax representation of the value
	// %v: the value in a default format

	// Print slices in go-syntax representation, default format, and length, capacity.
	fmt.Printf("slice1: %#v, %v, len[%d], cap[%d]\n", slice1, slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %#v, %v, len[%d], cap[%d]\n", slice2, slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3: %#v, %v, len[%d], cap[%d]\n", slice3, slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4: %#v, %v, len[%d], cap[%d]\n", slice4, slice4, len(slice4), cap(slice4))
	fmt.Printf("slice5: %#v, %v, len[%d], cap[%d]\n", slice5, slice5, len(slice5), cap(slice5))

	// Check if slice1 is nil.
	fmt.Printf("slice1 is nil: %t\n", slice1 == nil)
	// Check if slice2 is nil.
	fmt.Printf("slice2 is nil: %t\n", slice2 == nil)
}
