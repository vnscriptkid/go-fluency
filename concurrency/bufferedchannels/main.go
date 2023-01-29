package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func listenForChannel(channel chan int) {
	for {
		val := <-channel
		fmt.Printf("\nReceived value %v from channel", val)

		fmt.Println("\n...heavy work for 1 sec...")
		time.Sleep(1 * time.Second)
		fmt.Println("\n...DONE heavy work...")
	}
}

func listenForBufferedChannel(bufferedChannel chan int) {
	for {
		val := <-bufferedChannel
		color.Cyan("\nReceived value %v from channel", val)

		color.Cyan("\n...heavy work for 1 sec...")
		time.Sleep(1 * time.Second)
		color.Cyan("\n...DONE heavy work...")
	}
}

func pushToChannel(channel chan int) {
	for i := 1; i <= 100; i++ {
		// Pushing into unbufferedChannel
		fmt.Println("\tSending ", i, " to unbufferedChannel...")
		// Pushing to unbuffered channel can be blocked here as long as the goroutine is busy processing
		channel <- i
		fmt.Println("\tSent ", i, " to unbufferedChannel")
	}
}

func main() {
	unbufferedChannel := make(chan int)
	bufferedChannel := make(chan int, 5)

	go listenForChannel(unbufferedChannel)
	go listenForBufferedChannel(bufferedChannel)

	go pushToChannel(unbufferedChannel)

	for i := 1; i <= 100; i++ {
		// Pushing into unbufferedChannel
		color.Cyan("\tSending %v to unbufferedChannel...", i)
		// Pushing to unbuffered channel can be blocked here as long as the goroutine is busy processing
		bufferedChannel <- i
		color.Cyan("\tSent %v to unbufferedChannel", i)
	}

	fmt.Println("DONE")
	close(unbufferedChannel)
}
