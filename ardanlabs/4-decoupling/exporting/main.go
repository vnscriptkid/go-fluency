package main

import (
	"fmt"

	"github.com/vnscriptkid/go-fluency/ardanlabs/4-decoupling/exporting/animals"
	"github.com/vnscriptkid/go-fluency/ardanlabs/4-decoupling/exporting/counters"
	"github.com/vnscriptkid/go-fluency/ardanlabs/4-decoupling/exporting/persons"
	"github.com/vnscriptkid/go-fluency/ardanlabs/4-decoupling/exporting/users"
)

func main() {
	c := counters.New(10)

	fmt.Printf("Counter: %d\n", c)

	// animal type is not exported,
	// New is trying to return a value of unexported type
	// this pattern should be avoided
	a := animals.New()

	fmt.Printf("\nAnimal: %#v\n", a)
	fmt.Printf("Name: %s\n", a.Name)
	fmt.Printf("Age: %d\n", a.Age)

	u := users.New()
	fmt.Printf("\nUser: %#v\n", u)
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Email: %s\n", u.Email)
	fmt.Printf("IsMarried: %t\n", u.IsMarried())

	m := persons.Manager{
		Title: "Dev Manager",
		// Cannot construct a value of unexported type `person`
		// Solutions:
		// 1. Export the `person` type
		// 2. Use a factory function to construct a `person` value inside package `persons`
	}
	fmt.Printf("\nManager: %#v\n", m)

	m2 := persons.NewManager("Rel Manager 2", "John Doe", 123)
	fmt.Printf("\nManager2: %#v\n", m2)
	fmt.Printf("Name: %s\n", m2.Name)
	fmt.Printf("ID: %d\n", m2.ID)
	fmt.Printf("Title: %s\n", m2.Title)
}
