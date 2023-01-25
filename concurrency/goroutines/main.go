package main

import (
	"fmt"
	"time"
)

func printSth(content string) {
	fmt.Println(content)
}

func main() {
	go printSth("first line")

	printSth("second line")

	time.Sleep(1 * time.Second)
}
