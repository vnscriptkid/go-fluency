package main

import "fmt"

type notifier interface {
	notify()
}

// user implements the notifier interface
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// duration implements the notifier interface
type duration int

func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

// polymorphic function: asking for any concrete data that implements the notifier interface
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Bill", "bill@email.com"}
	// cannot use u (variable of type user) as notifier value in argument to sendNotification: user does not implement notifier (method notify has pointer receiver)
	// Because we're passing a value of type user, only methods with value receivers are recognized.
	// The notify method has a pointer receiver, so it's not available.
	// sendNotification(u)

	// To fix this, we can pass the address (pointer to user)
	// So, the notify method is available
	sendNotification(&u)

	// invalid operation: cannot take address of duration(42) (constant 42 of type duration)
	// sendNotification(&duration(42))

	// cannot call pointer method notify on duration
	// duration(42).notify()

	// Not all values are addressable (e.g., constants or intermediate expressions).
	// Methods with pointer receivers require an address to operate on, which these values lack.
}
