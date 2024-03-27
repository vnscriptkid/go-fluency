package main

import (
	"fmt"
)

func main() {
	type Person struct {
		a int
		b string
		c float64
		d bool
	}

	var a int
	var b string
	var c float64
	var d bool
	var pPerson *Person
	var person Person

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)
	fmt.Printf("var pPerson *Person \t %T [%v]\n", pPerson, pPerson)
	fmt.Printf("var person Person \t %T [%+v]\n", person, person)
	fmt.Printf("var person.a \t %T [%v]\n", person.a, person.a)
	fmt.Printf("var person.b \t %T [%v]\n", person.b, person.b)
	fmt.Printf("var person.c \t %T [%v]\n", person.c, person.c)
	fmt.Printf("var person.d \t %T [%v]\n", person.d, person.d)
}
