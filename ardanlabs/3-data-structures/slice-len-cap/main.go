package main

func main() {
	// slice := make([]string, 5)
	// slice[0] = "Apple"
	// slice[1] = "Orange"
	// slice[2] = "Banana"
	// slice[3] = "Grape"
	// slice[4] = "Plum"
	// slice[5] = "Runtime error"

	slice2 := make([]string, 5, 8)
	slice2[0] = "Apple"
	slice2[1] = "Orange"
	slice2[2] = "Banana"
	slice2[3] = "Grape"
	slice2[4] = "Plum"
	// slice2[5] = "Runtime error"
	// panic: runtime error: index out of range [5] with length 5
	// fmt.Printf("%v", slice2[5])
	// panic: runtime error: index out of range [5] with length 5
}
