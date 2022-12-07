package main

import "fmt"

func main() {
	_2darr := make([][]int, 3)

	fmt.Println("Empty _2darr: ", _2darr)

	for i := range _2darr {
		_2darr[i] = make([]int, 3)
	}

	fmt.Println("Arr after: ", _2darr)
}
