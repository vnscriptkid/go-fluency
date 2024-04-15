package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Polymorphism pattern 1
func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	makeItSpeak(dog)
	makeItSpeak(cat)

	// Polymorphism pattern 2
	animals := []Speaker{dog, cat}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
