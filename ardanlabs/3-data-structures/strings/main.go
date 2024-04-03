package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var strings [5]string

	strings[0] = "cat"

	fmt.Println(strings[0])
	fmt.Println(strings[1])
	fmt.Println(strings[2])

	fmt.Println("=====")

	// Define a string
	originalString := "Hello, Go!"

	// Slice the string
	slicedString := originalString[7:]

	// Use reflection to obtain the string header info
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&originalString))
	slicedStringHeader := (*reflect.StringHeader)(unsafe.Pointer(&slicedString))

	fmt.Printf("Original string: %s\n", originalString)
	fmt.Printf("Sliced string: %s\n\n", slicedString)

	fmt.Printf("Original String Header:\n  Data pointer: %v\n  Length: %v\n\n",
		stringHeader.Data, stringHeader.Len)

	fmt.Printf("Sliced String Header:\n  Data pointer: %v\n  Length: %v\n",
		slicedStringHeader.Data, slicedStringHeader.Len)

	fmt.Printf("Pointer difference: %v bytes\n",
		slicedStringHeader.Data-stringHeader.Data)
}
