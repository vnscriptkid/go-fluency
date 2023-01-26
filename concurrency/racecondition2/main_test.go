package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdout := os.Stdout

	r, w, err := os.Pipe()

	if err != nil {
		t.Errorf("err creating pipe")
	}

	os.Stdout = w

	main()

	err = w.Close()

	if err != nil {
		t.Errorf("err closing w pipe")
	}

	bSlice, err := io.ReadAll(r)

	if err != nil {
		t.Errorf("err reading r pipe")
	}

	output := string(bSlice)

	if !strings.Contains(output, "Final balance is 1000") {
		t.Errorf("Wrong result")
	}

	os.Stdout = stdout
}
