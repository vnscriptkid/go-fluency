package main

import "fmt"

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}

func main() {
	// Using copy
	slice1 := []string{"A", "B", "C", "D", "E"}
	slice2 := make([]string, len(slice1))

	// slice2 := make([]string, 4)
	// slice2 []string{"A", "B", "C", "D}

	// slice2 := make([]string, 6)
	// slice2 []string{"A", "B", "C", "D", "E", ""}

	copy(slice2, slice1)

	inspectSlice(slice1)
	inspectSlice(slice2)
}
