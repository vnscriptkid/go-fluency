package main

import "fmt"

// Structs are value types, so they are passed by value.

type MyStruct struct {
	Data []int
}

func main() {
	s := MyStruct{Data: []int{1, 2, 3}}
	modifyStruct(s)
	fmt.Println(s.Data) // [1 2 3], unchanged if modifyStruct() attempted to alter the slice's structure

	modifyStruct2(s)
	fmt.Println(s.Data) // [1 2 3], unchanged if modifyStruct2() attempted to alter the slice's structure

	pointerS := &MyStruct{Data: []int{1, 2, 3}}

	modifyStruct3(pointerS)
	fmt.Println(pointerS.Data) // [999 2 3], changed if modifyStruct3() attempted to alter the slice's elements

	modifyStruct4(pointerS)
	fmt.Println(pointerS.Data) // [999 2 3 4], changed if modifyStruct4() attempted to alter the slice's structure
}

func modifyStruct(m MyStruct) {
	m.Data[0] = 999
}

func modifyStruct2(m MyStruct) {
	m.Data = append(m.Data, 4) // This doesn't change the original slice in `s`
}

func modifyStruct3(m *MyStruct) {
	m.Data[0] = 999
}

func modifyStruct4(m *MyStruct) {
	m.Data = append(m.Data, 4) // This changes the original slice in `s`
}
