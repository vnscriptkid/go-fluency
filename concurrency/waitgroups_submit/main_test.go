package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_printSth(t *testing.T) {
	stdout := os.Stdout
	r, w, err := os.Pipe()

	if err != nil {
		t.Errorf("err creating pipe")
	}

	os.Stdout = w

	str = "example"

	printSth()

	// Important: Close before ReadAll
	err = w.Close()

	if err != nil {
		t.Errorf("err closing w pipe")
	}

	bSlice, err := io.ReadAll(r)

	if err != nil {
		t.Errorf("err reading r pipe")
	}

	actualOutput := string(bSlice)

	os.Stdout = stdout

	if !strings.Contains(actualOutput, "example") {
		t.Errorf("print wrong thing")
	}
}
