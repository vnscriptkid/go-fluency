package main

import (
	"fmt"
	"unsafe"
)

type reader interface {
	read(b []byte) (int, error)
}

// file implements the reader interface
type file struct {
	name  string
	size  int
	valid bool
}

func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// pipe implements the reader interface
type pipe struct {
	name string
}

func (pipe) read(b []byte) (int, error) {
	s := `{name: "Bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

// polymorphic function: asking for a concrete data that implements the reader interface
func retrieve(r reader) {
	// Print size and type of argument r
	fmt.Printf("[Type] %T [Size] %d bytes\n", r, unsafe.Sizeof(r))
	// interface values are represented as a two-word pair (16 bytes on 64-bit architectures)
	// TYPE: a pointer to information about the type stored in the interface value
	// VALUE: a pointer to the associated data

	switch r.(type) {
	case file:
		fmt.Println("File type")
	case pipe:
		fmt.Println("Pipe type")
	default:
		fmt.Println("Unknown type")
	}

	// cast the interface value to a concrete type
	// r.(file) or r.(pipe) will panic if the type assertion fails
	x, ok := r.(file)
	fmt.Printf("File: %+v, %v\n", x, ok)

	r.read([]byte{})
}

func main() {
	r := file{name: "data.json"}
	retrieve(r)

	r2 := pipe{"data.json"}
	retrieve(r2)
}
