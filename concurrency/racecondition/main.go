package main

import (
	"fmt"
	"sync"
)

var str string

var wg sync.WaitGroup

func updateString(newStr string, mutex *sync.Mutex) {
	defer wg.Done()

	mutex.Lock()
	str = newStr
	mutex.Unlock()
}

// go run -race ./
func main() {
	var m sync.Mutex

	wg.Add(2)

	go updateString("one", &m)
	go updateString("two", &m)

	wg.Wait()

	fmt.Println("String now: " + str)
}
