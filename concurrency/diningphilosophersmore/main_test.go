package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	testcases := []struct {
		name     string
		duration time.Duration
	}{
		{"zero sec", 0 * time.Second},
		{"one tenth sec", 100 * time.Microsecond},
		{"half sec", 500 * time.Millisecond},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			eatTime = tc.duration
			sleepTime = tc.duration
			thinkTime = tc.duration

			ordering = []string{}

			dine()

			if len(ordering) != 5 {
				t.Errorf("Expected 5 but got %v", len(ordering))
			}
		})
	}
}
