package main

import "fmt"

type Person struct {
	name string
}

const x = 99999999999999999999999999999999 // max: 256bits
const y = 12.34
const z = 3 * 0.333
const t = 3 / 2
const m = int8(8) * 1

// const n = int8(8) * 20;
// both is converted to same type
// int8 vs int => int8 vs int8
// float vs int => float vs float

// const m uint8 = 1000;

// const p = &Person {
// 	name: "",
// }

func main() {
	fmt.Printf("x is so big: %v", x > 0)
}
