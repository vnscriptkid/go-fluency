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

func Test_updateString(t *testing.T) {
	var wg sync.WaitGroup

	str = "old stuff"

	wg.Add(1)

	newStr := "new stuff"

	updateString(newStr, &wg)

	wg.Wait()

	if str != newStr {
		t.Errorf("Expected %s but got %s", newStr, str)
	}
}

func Test_main(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()

	bSlice, err := io.ReadAll(r)

	if err != nil {
		t.Errorf("err reading r pipe")
	}

	output := string(bSlice)

	items := strings.Split(output, "\n")

	if !strings.Contains(items[0], "one") {
		t.Errorf("Expected %s but got %s", "one", items[0])
	}

	if !strings.Contains(items[1], "two") {
		t.Errorf("Expected %s but got %s", "two", items[1])
	}

	if !strings.Contains(items[2], "three") {
		t.Errorf("Expected %s but got %s", "three", items[2])
	}

	os.Stdout = stdout
}
