package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("start goroutine")
		<-time.After(2 * time.Second)
		fmt.Println("line after")
	}()

	time.Sleep(5 * time.Second)
}
