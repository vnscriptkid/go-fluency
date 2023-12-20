package main

import (
	"fmt"

	"github.com/facebookgo/inject"
)

func main() {
	g := inject.Graph{}

	fmt.Println("hello", g)
}
