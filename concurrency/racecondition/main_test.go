package main

import (
	"sync"
	"testing"
)

func Test_updateString(t *testing.T) {
	str = "hello"

	var m sync.Mutex

	wg.Add(2)
	go updateString("first", &m)
	go updateString("second", &m)
	wg.Wait()

	if str != "first" && str != "second" {
		t.Errorf("Wrong result")
	}
}
