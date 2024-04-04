package main

import "encoding/binary"

func main() {
	x := []byte{0x0A, 0x15, 0x0e, 0x28, 0x05, 0x96, 0x0b, 0xd0, 0x0}
	a := x[0]
	b := binary.LittleEndian.Uint16(x[1:3])
	// binary.LittleEndian.Uint16(x[1:3]): x[1] + x[2] << 8
	// Explanation: x[1] is the least significant byte and x[2] is the most significant byte.
	// What does this do? It converts a 2-byte slice into a 16-bit integer.
	// How? It takes the least significant byte and adds it to the most significant byte shifted 8 bits to the left.
	// 0x0A: is 10 in decimal, size is 1 byte
	// 0x15: is 21 in decimal, size is 1 byte
	c := binary.LittleEndian.Uint16(x[3:5])
	// binary.LittleEndian.Uint16(x[3:5]): x[3] + x[4] << 8
	d := binary.LittleEndian.Uint32(x[5:9])
	println(a, b, c, d)
}
