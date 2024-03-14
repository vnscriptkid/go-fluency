package main

import (
	"0003/lib"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var name string
	flag.StringVar(&name, "n", "", "name of command")
	flag.Parse()

	name = strings.TrimSpace(name)

	if name == "" {
		panic("-n is required")
	}

	var result string
	var err error
	switch name {
	case "ls":
		result, err = lib.ExecuteLs()
	case "pwd":
		result, err = lib.ExecutePwd()
	default:
		panic(fmt.Sprintf("unsupported command `%s`", name))
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

// ./util -n ls
// ./util -n pwd
