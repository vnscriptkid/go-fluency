package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	// Access by idx
	fmt.Println("Num at index 2: ", arr[2])

	// Len of array
	fmt.Println("Len of arr: ", len(arr))

	// Capacity of array
	// cap === len
	fmt.Println("Capacity of arr: ", cap(arr))

	// append(arr, 3) !! invalid

	// Loop through arr
	for i := 0; i < len(arr); i++ {
		fmt.Println("Ele at idx "+strconv.Itoa(i)+" is: ", arr[i])
	}
}
