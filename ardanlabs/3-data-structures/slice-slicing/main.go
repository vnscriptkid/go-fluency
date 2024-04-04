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
	slice1 := []string{"A", "B", "C", "D", "E", "F", "G"}
	slice2 := slice1[2:4]
	inspectSlice(slice1)
	inspectSlice(slice2)

	// Length[7] Capacity[7]
	// [0] 0x140000ba000 A
	// [1] 0x140000ba010 B
	// [2] 0x140000ba020 C
	// [3] 0x140000ba030 D
	// [4] 0x140000ba040 E
	// [5] 0x140000ba050 F
	// [6] 0x140000ba060 G
	// Length[2] Capacity[5]
	// [0] 0x140000ba020 C << Same address as slice1[2]
	// [1] 0x140000ba030 D << Same address as slice1[3]

	// SIDE EFFECT: Changing slice2 will change slice1
	slice2[0] = "CC"

	inspectSlice(slice1)
	inspectSlice(slice2)

	// 2 slices share the same backing array
}
