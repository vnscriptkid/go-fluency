package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSth(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	initialInput := "hello"

	printSth(initialInput, &wg)

	wg.Wait()

	err := w.Close()

	if err != nil {
		t.Errorf("err closing w pipe")
	}

	bSlice, err := io.ReadAll(r)

	if err != nil {
		t.Errorf("err reading from r pipe")
	}

	output := string(bSlice)

	os.Stdout = stdout

	if !strings.Contains(output, initialInput) {
		t.Errorf("expected %s but got %s", initialInput, output)
	}
}
