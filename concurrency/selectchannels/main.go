package main

import (
	"fmt"
	"time"
)

func server1(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "This is from server 1"
	}
}

func server2(c chan string) {
	for {
		time.Sleep(2 * time.Second)
		c <- "This is from server 2"
	}
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		fmt.Println()
		fmt.Println("Starting loop")

		// When multiple case match, one case will be picked randomly to run
		select {
		case res := <-channel1:
			fmt.Printf("Case 1: %s", res)
		case res := <-channel1:
			fmt.Printf("Case 2: %s", res)
		case res := <-channel2:
			fmt.Printf("\t\tCase 3: %s", res)
		case res := <-channel2:
			fmt.Printf("\t\tCase 4: %s", res)
		}
	}
}
