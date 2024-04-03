package main

import (
	"encoding/json"
	"fmt"
)

// call of Unmarshal passes non-pointer as second argument
// func decodeJSON(data string, target []int) error {
// 	// return json.Unmarshal([]byte(data), target)
// }

// func decodeJSON(data string, target interface{}) error {
// 	return json.Unmarshal([]byte(data), target)
// }

func decodeJSON(data string, target *[]int) error {
	return json.Unmarshal([]byte(data), target)
}

func main() {
	data := "[1,2,3,4]"
	var mySlice []int
	err := decodeJSON(data, &mySlice)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(mySlice) // Output: [1 2 3 4]
}
