package main

import "fmt"

func main() {
	// Append gets its own copy of a slice value, it mutates its own copy, then it returns a copy back to the caller.
	// Value semantics are used

	var data []string
	for record := 1; record <= 50; record++ {
		data = append(data, fmt.Sprintf("Rec: %d", record))
		fmt.Printf("Length[%d] Capacity[%d] First Element Address: %p\n", len(data), cap(data), &data[0])
	}

	// Length[1] Capacity[1] First Element Address: 0x14000010040
	// Length[2] Capacity[2] First Element Address: 0x1400010c000
	// Length[3] Capacity[4] First Element Address: 0x14000030080
	// Length[4] Capacity[4] First Element Address: 0x14000030080
	// Length[5] Capacity[8] First Element Address: 0x1400010e000
	// Length[6] Capacity[8] First Element Address: 0x1400010e000
	// Length[7] Capacity[8] First Element Address: 0x1400010e000
	// Length[8] Capacity[8] First Element Address: 0x1400010e000
	// Length[9] Capacity[16] First Element Address: 0x14000110000
	// Length[10] Capacity[16] First Element Address: 0x14000110000

	// Slice is doubling its capacity when it reaches its current capacity
	// But it still maintains contiguous memory
	// This is because the runtime needs to allocate a new array to store the new elements.
	// First Element Address is changing because the slice is being reallocated in memory.
}
