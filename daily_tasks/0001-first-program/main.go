package main

import (
	"fmt"
	"os"
)

type User struct {
	name string
	age  int
}

func main() {
	username := os.Getenv("USER")

	u := &User{
		name: username,
		age:  28,
	}

	users := []*User{u, {name: "someone", age: 30}}

	fmt.Printf("Hello %s\n", username)
	fmt.Printf("Here's details: %+v", u)
	fmt.Printf("Here's all users: %+v", users)
}
