package main

import "fmt"

// No embedding
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("[INNER] Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// Pointer semantics embedding
type admin struct {
	*user
	level string
}

func (a *admin) notify() {
	fmt.Printf("[OUTER] Sending Admin Email To %s<%s>\n",
		a.name,
		a.email)
}

type notifier interface {
	notify()
}

// Polymorphic function: it can accept any type that implements the notifier interface
// Behaviors of the function will change based on the type of the value that is passed in
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	////////////////////////////
	// Declare an admin user.
	////////////////////////////
	ad := admin{
		// When defining a struct type, still need to explicitly declare the type of the embedded struct
		user: &user{
			name:  "Jon Calhoun",
			email: "admin@gmail.com",
		},
		level: "super",
	}

	fmt.Printf("Admin: %#v\n", ad)
	fmt.Printf("Level: %#v\n", ad.level)
	// Access the fields of the embedded struct
	fmt.Printf("Name: %#v\n", ad.user.name)
	fmt.Printf("Email: %#v\n", ad.user.email)
	// Access the methods of the embedded struct
	ad.user.notify()
	// Access the fields of inner struct directly
	fmt.Printf("Name: %#v\n", ad.name)
	fmt.Printf("Email: %#v\n", ad.email)
	// Access the methods of inner struct directly
	ad.notify()

	// Does ad implement the notifier interface?
	// Yes, it does it directly, promotion is NOT needed in this case

	sendNotification(&ad)
	// cannot use ad (variable of type admin) as notifier value in argument to sendNotification:
	// admin does not implement notifier (method notify has pointer receiver)
	// sendNotification(ad)
}
