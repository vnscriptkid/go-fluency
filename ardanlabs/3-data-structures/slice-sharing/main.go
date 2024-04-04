package main

import "fmt"

// Slice is a reference to an underlying array.
// When you pass a slice to a function, you are passing a reference to the same underlying array.
// If you modify the slice in the function, you are modifying the same underlying array.
func appendToSlice(slice []int) []int {
	// This append could cause the slice to reallocate its underlying array
	// if the capacity is exceeded.
	return append(slice, 4)
}

type MySlice struct {
	Slice []int
}

func appendToMySlice(s *MySlice) {
	s.Slice = append(s.Slice, 4)
}

func main() {
	// CASE 1: Create a slice with a length of 3 and a capacity of 4.
	// Appending to this slice will not cause the slice to reallocate its
	// originalSlice := make([]int, 3, 4)
	// originalSlice[0] = 1
	// originalSlice[1] = 2
	// originalSlice[2] = 3

	// CASE 2:Create a slice with a length of 3 and a capacity of 3
	// Appending to this slice will cause the slice to reallocate its
	originalSlice := []int{1, 2, 3}
	fmt.Printf("[originalSlice] %+v [FirstAddr] %p\n", originalSlice, &originalSlice[0])

	// Append to the slice in a function.
	// WARNING: FORGETTING TO ASSIGN THE RETURN VALUE
	newSlice := appendToSlice(originalSlice)
	fmt.Printf("[newSlice] %+v [FirstAddr] %p\n", newSlice, &newSlice[0])

	// Share and modify the original slice.
	sharedSlice := originalSlice
	sharedSlice[0] = 10

	fmt.Printf("[sharedSlice] %+v [FirstAddr] %p\n", sharedSlice, &sharedSlice[0])
	fmt.Printf("[originalSlice] %+v [FirstAddr] %p\n", originalSlice, &originalSlice[0])

	mySlice := MySlice{Slice: []int{1, 2, 3}}
	fmt.Printf("[mySlice] %+v [FirstAddr] %p\n", mySlice.Slice, &mySlice.Slice[0])
	appendToMySlice(&mySlice)
	fmt.Printf("[mySlice] %+v [FirstAddr] %p\n", mySlice.Slice, &mySlice.Slice[0])
}
