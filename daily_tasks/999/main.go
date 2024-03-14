package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

func main() {
	fmt.Println("hi there")

	jsonStr := `{"name":"John","age":30}`

	person := Person{}
	json.Unmarshal([]byte(jsonStr), &person)

	// fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	fmt.Printf("Name: %s, Age: %d, Job: %s\n", person.Name, person.Age, person.Job)
}
