package main

import "fmt"

// Maps are reference types, so they are passed by reference.

func modifyMap(m map[string][]int) {
	m["key"][0] = 999
}

func modifyMap2(m map[string][]int) {
	m["key"] = append(m["key"], 4)
}

func main() {
	m := make(map[string][]int)
	m["key"] = []int{1, 2, 3}
	modifyMap(m)
	fmt.Println(m["key"]) // [999 2 3], if modifyMap directly modifies the slice's elements

	modifyMap2(m)
	fmt.Println(m["key"]) // [1 2 3 4], if modifyMap2 directly modifies the slice's structure
}
