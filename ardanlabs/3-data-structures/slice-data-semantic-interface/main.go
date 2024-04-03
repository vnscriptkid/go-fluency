package main

import "fmt"

func modifyInterface(i interface{}) {
	// update first element
	i.([]int)[0] = 999

	i = append(i.([]int), 4)
}

func main() {
	var i interface{} = []int{1, 2, 3}
	modifyInterface(i)

	fmt.Println(i) // Output: [999 2 3]
}
