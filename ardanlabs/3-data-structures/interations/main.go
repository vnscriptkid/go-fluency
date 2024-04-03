package main

import (
	"fmt"
	"time"
)

func main() {

	var numbers [10000000]int

	// Value Semantic Iteration
	start := time.Now()
	sum1 := 0
	for _, fruit := range numbers {
		sum1 += fruit
	}
	elapsed1 := time.Since(start)
	// Pointer Semantic Iteration
	start = time.Now()
	sum2 := 0
	for i := range numbers {
		sum2 += numbers[i]
	}
	elapsed2 := time.Since(start)

	println("Value Semantic Iteration Elapsed in Nanoseconds: ", elapsed1.Nanoseconds())
	println("Pointer Semantic Iteration Elapsed in Nanoseconds: ", elapsed2.Nanoseconds())
	fmt.Printf("Difference: %v\n", elapsed1.Nanoseconds()-elapsed2.Nanoseconds())
}
