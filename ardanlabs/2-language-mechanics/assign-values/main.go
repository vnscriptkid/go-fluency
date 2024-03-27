package main

import "fmt"

func main() {
	type Animal struct {
		age  int
		name string
	}

	type Cat struct {
		age  int
		name string
	}

	type Dog struct {
		age  int32
		name string
	}

	var chicken struct {
		age  int32
		name string
	}

	a := Animal{age: 5, name: "Fluffy"}
	c := Cat(a)
	var d Dog
	d = chicken

	fmt.Printf("a: %T [%+v]\n", a, a)
	fmt.Printf("c: %T [%+v]\n", c, c)
	fmt.Printf("d: %T [%+v]\n", d, d)
}
