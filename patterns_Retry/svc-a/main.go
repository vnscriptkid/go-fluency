package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	lib "github.com/vnscriptkid/go-fluency/patterns_Retry/shared/lib"
)

func exampleFunction() error {
	// Simulating a random error with 50% probability
	if rand.Float32() < 0.5 {
		return errors.New("example error")
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Choose your desired strategy
	// strategy := lib.FixedTimeoutStrategy{Timeout: 200 * time.Millisecond}
	// strategy := lib.IncrementalTimeoutStrategy{InitialTimeout: 200 * time.Millisecond, Increment: 100 * time.Millisecond}
	strategy := lib.ExponentialBackoffStrategy{InitialBackoff: 200 * time.Millisecond}

	err := lib.Retry(exampleFunction, 5, strategy)
	if err != nil {
		if errors.Is(err, lib.ErrMaxRetriesReached) {
			fmt.Println("Max retries reached. Giving up.")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
		return
	}

	fmt.Println("Success!")
}
