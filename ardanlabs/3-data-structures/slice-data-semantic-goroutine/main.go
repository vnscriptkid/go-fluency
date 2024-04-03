package main

import "fmt"

func main() {
	originalSlice := []int{1, 2, 3}
	ch := make(chan []int)
	go func() {
		ch <- originalSlice
	}()
	receivedSlice := <-ch

	// this will change the originalSlice as well
	receivedSlice[0] = 999

	// this will not change the originalSlice
	receivedSlice = append(receivedSlice, 4)

	fmt.Printf("originalSlice: %v\n", originalSlice)
	fmt.Printf("receivedSlice: %v\n", receivedSlice)

}
