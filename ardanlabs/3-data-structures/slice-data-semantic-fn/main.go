package main

import "fmt"

// Copy of the original slice is passed to the function including
// 1. the pointer to first element of the array
// 2. the length of the slice
// 3. the capacity of the slice
func modifySlice(s []int) {
	s[0] = 999 // This change will affect the original slice as well.

	// Try enlarging the slice.
	s = append(s, 4)
	fmt.Println(s) // Output: [999 2 3 4]
}

func main() {
	originalSlice := []int{1, 2, 3}
	modifySlice(originalSlice)
	fmt.Println(originalSlice) // Output: [999 2 3]
}
