package main

import (
	"fmt"
	"unsafe"
)

type printer interface {
	print()
}
type canon struct {
	name  string
	model string
	year  int
}

// value receiver: this method will operate on a copy of canon
func (c canon) print() {
	fmt.Printf("Printer Name: %s\n", c.name)
}

type epson struct {
	name  string
	model string
	year  int
}

// pointer receiver: this method will operate on the actual epson
func (e *epson) print() {
	fmt.Printf("Printer Name: %s\n", e.name)
}
func main() {
	c := canon{"PIXMA TR4520", "xxx", 2012}
	e := epson{"WorkForce Pro WF-3720", "yyy", 2020}
	printers := []printer{
		c, // Assuming this is a copy of c, like passing struct to a function by value
		&e,
	}
	c.name = "[NEW ]PROGRAF PRO-1000"
	e.name = "[NEW] Home XP-4100"

	// print address of each printer
	fmt.Printf("Address of canon printer: %p\n", &c)
	fmt.Printf("Address of epson printer: %p\n", &e)
	fmt.Println()

	for _, p := range printers {
		// print size of each printer
		fmt.Printf("Size of printer: %d\n", unsafe.Sizeof(p))

		interfaceHeader := (*[2]uintptr)(unsafe.Pointer(&p))
		fmt.Printf("Type pointer (itab): %x\n", interfaceHeader[0])
		fmt.Printf("Data pointer: %x\n", interfaceHeader[1])

		// Observation: The size of each printer is 16 bytes
		// This is because the interface value is a two-word data structure
		// 1. Type pointer (itab): The first word points to information about the type.
		// This includes a method table for the dynamic type,
		// which allows the correct method implementation to be called when a method is invoked on the interface value.
		// 2. Data pointer: The second word points to the actual data held by the variable.
		// This could be a value (for small types) or a pointer to the value (for larger types).

		// Polymorphism: The print method will be called based on the actual type of the printer
		// Behavior of the print method will be different for canon and epson printer
		p.print()
		// Only epson printer will be able to print the new name
		// This is because canon printer is passed by value
		// while epson printer is passed by reference
		fmt.Println()
	}
}
