package main

import "fmt"

func main() {
	slice := make([]string, 5, 8)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"
	inspectSlice(slice)

	// Keep appending
	slice = append(slice, "Kiwi")
	slice = append(slice, "Peach")
	slice = append(slice, "Strawberry")

	// Length[8] Capacity[8]
	inspectSlice(slice)

	// Keep appending beyond capacity
	// Length[11] Capacity[16]: Capacity is doubled. Why? Because the slice is full and the runtime needs to allocate a new array to store the new elements.
	slice = append(slice, "Mango")
	inspectSlice(slice)
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))

	// Only traverse the slice length
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}

// [0] 0x1400012c000 Apple
// [1] 0x1400012c010 Orange
// [2] 0x1400012c020 Banana
// [3] 0x1400012c030 Grape
// [4] 0x1400012c040 Plum

// Size of a string is 16 bytes (2 words of 8 bytes each)
