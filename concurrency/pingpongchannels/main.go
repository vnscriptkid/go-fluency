package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping

		pong <- strings.ToUpper(s)
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	for {
		fmt.Println("\nEnter something (q to quit): ")

		var userInput string
		fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		// send user input to ping channel
		fmt.Println("-- PING")
		ping <- userInput

		// receives response from pong channel (blocking op)
		res := <-pong
		fmt.Printf("-- PONG: %s", res)
	}

	fmt.Println("DONE")
	close(ping)
	close(pong)
}
