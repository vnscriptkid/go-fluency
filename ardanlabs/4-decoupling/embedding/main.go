package main

import "fmt"

// No embedding
type user struct {
	name  string
	email string
}
type admin struct {
	person user
	level  string
}

// Value semantics embedding
type user2 struct {
	name  string
	email string
}

func (u *user2) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

type admin2 struct {
	user2
	level string
}

// Pointer semantics embedding
type user3 struct {
	name  string
	email string
}

func (u *user3) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

type admin3 struct {
	// Embedding PROMOTES the fields and methods of the inner type to the outer type
	*user3
	level string
}

func main() {
	////////////////////////////
	// Declare an admin user.
	////////////////////////////
	ad := admin{
		person: user{
			name:  "Jon Calhoun",
			email: "admin@gmail.com",
		},
		level: "super",
	}

	fmt.Printf("\nAdmin1: %#v\n", ad)
	fmt.Printf("Name: %#v\n", ad.person.name)
	fmt.Printf("Email: %#v\n", ad.person.email)
	fmt.Printf("Level: %#v\n", ad.level)

	////////////////////////////
	// Declare an admin user.
	////////////////////////////
	ad2 := admin2{
		user2: user2{
			name:  "Jon Calhoun",
			email: "user2@gmail.com",
		},
		level: "super",
	}

	fmt.Printf("\nAdmin2: %#v\n", ad2)
	fmt.Printf("Level: %#v\n", ad2.level)
	// Can access the fields of user2 directly from outer struct admin2
	fmt.Printf("Name: %#v\n", ad2.name)
	fmt.Printf("Email: %#v\n", ad2.email)
	// Can still access the fields step by step
	fmt.Printf("Name: %#v\n", ad2.user2.name)
	fmt.Printf("Email: %#v\n", ad2.user2.email)
	// Call the method of the inner type
	ad2.notify()

	////////////////////////////
	// Declare an admin user.
	////////////////////////////
	ad3 := admin3{
		user3: &user3{
			name:  "Jon Calhoun",
			email: "user3@gmail.com",
		},
		level: "super",
	}

	fmt.Printf("\nAdmin3: %#v\n", ad3)
	fmt.Printf("Level: %#v\n", ad3.level)
	// Can access the fields of user3 directly from outer struct admin3
	fmt.Printf("Name: %#v\n", ad3.name)
	fmt.Printf("Email: %#v\n", ad3.email)
	// Can still access the fields step by step
	fmt.Printf("Name: %#v\n", ad3.user3.name)
	fmt.Printf("Email: %#v\n", ad3.user3.email)
	// Call the method of the inner type
	ad3.notify()

}
