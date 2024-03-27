package main

import "fmt"

func main() {
	x := 42.3
	y := float32(x)
	z := int(y)

	fmt.Printf("x: %T [%v]\n", x, x)
	fmt.Printf("y: %T [%v]\n", y, y)
	fmt.Printf("z: %T [%v]\n", z, z)
}
