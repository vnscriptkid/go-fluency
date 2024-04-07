package main

import (
	"fmt"
	"os"
	"strings"
)

type stringsList []string

// Generally, for internal types like: slice, map, channel, function, interface
// Use value receivers
func (s stringsList) join(sep string) string {
	return strings.Join(s, sep)
}

func main() {
	s := stringsList{"hello", "world"}
	s.join(" ")

	// encoding.TextUnmarshaler
	// type TextUnmarshaler interface {
	// 	UnmarshalText(text []byte) error
	// }

	// func (t *Time) UnmarshalText(data []byte) error {
	// 	var err error
	// 	*t, err = parseStrictRFC3339(data)
	// 	return err
	// }

	// open a file
	f, _ := os.Open("file.txt")

	fmt.Println(f)
}
