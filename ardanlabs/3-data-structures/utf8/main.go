package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "世界 means world"

	fmt.Printf("String: %q\n", s[:])

	// size of codepoint ranges from 1 to 4 bytes (int32)
	// chinese characters are 3 bytes
	var buf [utf8.UTFMax]byte

	// Iterate over the string, move codepoint by codepoint
	// rune is an alias for int32 (represents a Unicode code point)
	for i, r := range s {
		rl := utf8.RuneLen(r)
		si := i + rl

		// buf[:] means the entire array
		copy(buf[:], s[i:si])

		fmt.Printf("%2d: %q; codepoint: %#6x; runelen: %v bytes; encoded bytes: %#v\n",
			i, r, r, rl, buf[:rl])
	}
}
