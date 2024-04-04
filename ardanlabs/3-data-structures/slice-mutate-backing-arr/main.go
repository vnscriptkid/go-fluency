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

func sideEffect() {
	slice1 := []string{"A", "B", "C", "D", "E"}
	slice2 := slice1[2:4]

	// SIDE EFFECT: Changing slice2 will change slice1
	slice2 = append(slice2, "CHANGED")
	inspectSlice(slice1)
	inspectSlice(slice2)
}

func noSideEffect() {
	slice1 := []string{"A", "B", "C", "D", "E"}
	slice2 := slice1[2:4:4]
	// slice2: Length[2] Capacity[2]

	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice2), cap(slice2))

	// NO SIDE EFFECT: Changing slice2 will not change slice1
	slice2 = append(slice2, "CHANGED")
	inspectSlice(slice1)
	inspectSlice(slice2)
}

func main() {
	fmt.Println("Side effect")
	sideEffect()

	fmt.Println("No side effect")
	noSideEffect()
	// Length[5] Capacity[5]
	// [0] 0x1400010e0a0 A
	// [1] 0x1400010e0b0 B
	// [2] 0x1400010e0c0 C
	// [3] 0x1400010e0d0 D
	// [4] 0x1400010e0e0 E <<<< Does not change
	// Length[3] Capacity[4]
	// [0] 0x1400012a040 C <<<< New backing array
	// [1] 0x1400012a050 D
	// [2] 0x1400012a060 CHANGED
}
