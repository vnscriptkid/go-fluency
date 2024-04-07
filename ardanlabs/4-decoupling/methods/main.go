package main

import "fmt"

type user struct {
	name  string
	email string
}

// Should not mix value and pointer receivers
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}
func (u *user) changeEmail(email string) {
	u.email = email
	fmt.Printf("Changed User Email To %s\n", email)
}

func main() {
	user1 := user{
		name:  "Bill",
		email: "bill@test.com",
	}

	user1.notify()
	user1.changeEmail("bill@new.com")

	fmt.Printf("User1: %#v\n", user1)

	fmt.Println()
	user2 := &user{
		name:  "Lisa",
		email: "lisa@test.com",
	}

	user2.notify()
	user2.changeEmail("lisa@new.com")

	fmt.Printf("User2: %#v\n", user2)
}
