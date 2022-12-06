package main

import "fmt"

func powerSeries1(x int) (int, int) {
	return x * x, x * x * x
}

func powerSeries2(x int) (square int, cube int) {
	square = x * x
	cube = x * x * x

	// This implicitly returns a tuple of 2 eles: square and cube
	return
}

func main() {
	square1, cube1 := powerSeries1(2)

	fmt.Println("Square 1: ", square1, "Cube 1: ", cube1)

	square2, cube2 := powerSeries2(3)

	fmt.Println("Square 2: ", square2, "Cube 2: ", cube2)
}
