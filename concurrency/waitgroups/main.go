package main

import (
	"fmt"
	"sync"
)

func printSth(content string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(content)
}

func main() {
	var wg sync.WaitGroup

	fruits := []string{
		"apple",
		"orange",
		"banana",
		"grape",
		"avocado",
		"pineapple",
		"kiwi",
		"coconut",
		"jackfruit",
		"durian",
	}

	wg.Add(len(fruits))

	for i, fruit := range fruits {
		go printSth(fmt.Sprintf("%d : %s", i, fruit), &wg)
	}

	// The WaitGroup is used to wait for a collection of goroutines to finish executing
	// before continuing on with the rest of the program
	wg.Wait()

	wg.Add(1)
	printSth("last line to be ran", &wg)
}
