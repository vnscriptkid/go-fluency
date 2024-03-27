package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a int32
		b int16
		c bool
		d bool
	}

	var y struct {
		c bool
		b int16
		d bool
		a int32
	}

	// Size of x
	fmt.Printf("Size of x: %d bytes\n", unsafe.Sizeof(x))
	fmt.Printf("Size of y: %d bytes\n", unsafe.Sizeof(y))
}
