package main

import (
	"fmt"
	"sync"
)

var str string

func updateString(newStr string, wg *sync.WaitGroup) {
	defer wg.Done()
	str = newStr
}

func printSth() {
	fmt.Println(str)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go updateString("one", &wg)
	wg.Wait()
	printSth()

	wg.Add(1)
	go updateString("two", &wg)
	wg.Wait()
	printSth()

	wg.Add(1)
	go updateString("three", &wg)
	wg.Wait()
	printSth()
}
