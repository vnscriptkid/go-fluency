package main

import "fmt"

func main() {
	// make a channel of int values
	ch := make(chan int)

	// make a goroutine that loops 1 -> 5, passing cur value to channel
	// at the end of loop, close the channel

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// infinite loop
	// decide what to do given signal from channel (select-case)
	// break loop if channel is closed
	for {
		select {
		// when a channel is closed, any attempt to read from it
		// will return the zero value of the channel's type
		// along with a boolean value indicating whether the channel is closed or not
		case i, ok := <-ch:
			if !ok {
				fmt.Printf("Channel is closed i: %d, ok: %v", i, ok)
				return
			}
			fmt.Printf("-- Consuming value %d", i)
			fmt.Println()
		}
	}
}
