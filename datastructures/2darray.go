package main

import "fmt"

func main() {
	var arr [2][2]int

	fmt.Println("Empty arr: ", arr)

	arr[0][0] = 1
	arr[0][1] = 2
	arr[1][0] = 3
	arr[1][1] = 4

	fmt.Println("Arr after assigning values: ", arr)
}
