package main

import (
	"fmt"
)

var str string

func updateString(newStr string) {
	str = newStr
}

func printSth() {
	fmt.Println(str)
}

func main() {

	updateString("one")
	printSth()

	updateString("two")
	printSth()

	updateString("three")
	printSth()
}
