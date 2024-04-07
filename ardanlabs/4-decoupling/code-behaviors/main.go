package main

import "fmt"

type data struct {
	name string
	age  int
}

// Value semantics
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// Pointer semantics
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {
	d := data{
		name: "Bill",
	}
	// Method is just a function with a receiver argument
	f1 := d.displayName

	f1()
	d.name = "Joan"
	// f1 = d.displayName
	f1()

	// Why second call to f1() prints "Bill" instead of "Joan"?
	// Because the receiver is passed by value, so the receiver is copied
	// And receiver is set to "Bill" when f1 is assigned to d.displayName

	fmt.Println()
	d2 := data{
		name: "Bill",
	}
	f2 := d2.setAge
	f2(45)
	d2.name = "Sammy"
	f2(45)
}
