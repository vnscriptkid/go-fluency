package main

import "fmt"

type user struct {
	name string
	age  int
}

func modifyUserDangerously1(users []user) {
	for _, user := range users {
		user.age += 1
	}
}

func modifyUserDangerously2(users []user) (users2 []user) {
	users = append(users, user{"K", 60})

	return users

}

func modifyUserSafely2(users []*user) (users2 []*user) {
	users = append(users, &user{"N", 60})
	users = append(users, &user{"O", 70})
	users = append(users, &user{"P", 80})
	users = append(users, &user{"Q", 90})
	users = append(users, &user{"R", 100})
	users = append(users, &user{"S", 110})
	users = append(users, &user{"T", 120})
	users = append(users, &user{"U", 130})

	return users
}

func modifyUserSafely(users []*user) {
	for _, user := range users {
		user.age += 1
	}
}

func main() {
	//////////////////////
	//////////////////////
	users1 := []user{
		{"A", 30},
	}
	modifyUserDangerously1(users1)
	fmt.Printf("Users1: %+v\n", users1)

	//////////////////////
	//////////////////////
	users2 := []user{
		{"T", 30},
	}
	userT := &users2[0]
	userT.age += 1
	fmt.Printf("Users2: %+v\n", users2)
	users2 = modifyUserDangerously2(users2)
	userT.age += 100
	fmt.Printf("Users2: %+v\n", users2)

	//////////////////////
	//////////////////////
	fmt.Println(`//////////////////////`)
	listOfPointers := []*user{
		{"X", 40},
		{"Y", 50},
	}
	modifyUserSafely(listOfPointers)
	fmt.Printf("Users: %+v, %+v\n", listOfPointers[0], listOfPointers[1])

	//////////////////////
	//////////////////////
	fmt.Println(`//////////////////////`)
	userM := &user{"M", 40}
	listOfPointers2 := []*user{
		userM,
	}
	// Length and Capacity of listOfPointers2
	fmt.Printf("Length[%d] Capacity[%d] Slice [%+v]\n", len(listOfPointers2), cap(listOfPointers2), listOfPointers2)

	fmt.Println()
	listOfPointers2 = modifyUserSafely2(listOfPointers2)
	fmt.Printf("Length[%d] Capacity[%d] Slice [%+v]\n", len(listOfPointers2), cap(listOfPointers2), listOfPointers2)
	userM.age += 1
	fmt.Printf("Users: %+v, %+v\n", listOfPointers2[0], listOfPointers2[1])

}
