package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	var balance int

	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go func(i int) {
			defer wg.Done()

			m.Lock()
			temp := balance
			newBalance := temp + 1
			balance = newBalance
			m.Unlock()

			fmt.Printf("At loop %d balance get updated %d => %d", i, temp, newBalance)
			fmt.Println()
		}(i)
	}
	wg.Wait()

	fmt.Printf("Final balance is %d", balance)
}
