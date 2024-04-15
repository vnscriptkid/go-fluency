package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Describe() {
	fmt.Println("Animal:", a.Name)
}

type Dog struct {
	Animal // This is composition; Dog "has-a" Animal
	Breed  string
}

func (d Dog) Describe() {
	fmt.Println("Dog:", d.Name, "Breed:", d.Breed)
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Max"},
		Breed:  "Golden Retriever",
	}
	d.Describe() // Calls Dog's Describe, demonstrating polymorphism
}
