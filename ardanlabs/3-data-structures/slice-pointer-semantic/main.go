package main

import "fmt"

func enlargeSlice(pSlice *[]int) {
	// pSlice[0] = 999
	// invalid operation: cannot index pSlice (variable of type *[]int)

	// Modifying the slice's first element.
	(*pSlice)[0] = 999

	// Appending an element, which might change the slice's length and capacity.
	*pSlice = append(*pSlice, 4)
}

func main() {
	originalSlice := []int{1, 2, 3}
	enlargeSlice(&originalSlice)
	fmt.Println(originalSlice) // Output: [1 2 3 4]
}
