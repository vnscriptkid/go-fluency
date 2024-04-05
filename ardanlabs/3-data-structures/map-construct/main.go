package main

import "fmt"

type user struct {
	name     string
	username string
}

func main() {
	// ### 1: ZERO VALUE
	// Construct a map set to its zero value,
	// that can store user values based on a key of type string.
	// Trying to use this map will result in a runtime error (panic).
	var users map[string]user
	fmt.Printf("users: %#v\n", users) // map[string]main.user(nil)
	fmt.Printf("users == nil: %v\n", users == nil)
	fmt.Printf("len(users): %v\n", len(users))
	fmt.Printf("value of users[\"key\"]: %#v\n", users["key"])
	// users["key"] = user{name: "John Doe", username: "johndoe"}
	// >>>>>>panic: assignment to entry in nil map
	fmt.Println()

	// ### 2: MAKE FUNCTION
	// Construct a map initialized using make,
	// that can store user values based on a key of type string.
	users2 := make(map[string]user, 0)
	fmt.Printf("users2: %#v\n", users2)
	fmt.Printf("users2 == nil: %v\n", users2 == nil)
	fmt.Printf("len(users2): %v\n", len(users2))
	fmt.Printf("value of users2[\"key\"]: %#v\n", users2["key"])
	users2["key"] = user{name: "John Doe", username: "johndoe"}
	fmt.Printf("users2: %#v\n", users2)
	fmt.Printf("len(users2): %v\n", len(users2))

	fmt.Println()

	// ### 3: EMPTY LITERAL CONSTRUCTION
	// Construct a map initialized using empty literal construction,
	// that can store user values based on a key of type string.
	users3 := map[string]user{}
	fmt.Printf("users3: %#v\n", users3)
	fmt.Printf("users3 == nil: %v\n", users3 == nil)
	fmt.Printf("len(users3): %v\n", len(users3))
	fmt.Printf("value of users3[\"key\"]: %#v\n", users3["key"])
	users3["john"] = user{name: "John Doe", username: "johndoe"}
	users3["jane"] = user{name: "Jane Doe", username: "janedoe"}
	users3["smith"] = user{name: "John Smith", username: "johnsmith"}
	fmt.Printf("users3: %#v\n", users3)

	// Difference between 2 and 3:
	// 2: make function is used to create a map with an initial capacity.
	// 3: empty literal construction is used to create a map with no initial capacity.

	// Iterating over a map:
	// The order of iteration is not guaranteed to be the same as the order of insertion.
	for key, value := range users3 {
		fmt.Printf("key: %v, value: %#v\n", key, value)
	}

	// Check if a key exists in a map:
	// ALWAYS use the comma ok idiom to check if a key exists in a map.
	// DO NOT use the zero value of the value type to check if a key exists in a map.
	val, exists := users3["xxx"]
	fmt.Printf("val: %#v, exists: %v\n", val, exists)

	// Delete a key from a map:
	fmt.Println()
	delete(users3, "john")
	delete(users3, "xxx") // no error if key does not exist
	delete(users3, "jane")
	fmt.Printf("users3: %#v\n", users3)
	fmt.Printf("len(users3): %v\n", len(users3))

	// Key map restrictions:
	// INVALID: type slice []user
	type slice [1]user
	users4 := map[slice]string{}
	fmt.Printf("users4: %#v\n", users4)
}
