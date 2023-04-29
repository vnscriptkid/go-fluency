package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	shared "github.com/vnscriptkid/go-fluency/patterns_Retry/shared"
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
	// strategy := shared.FixedTimeoutStrategy{Timeout: 200 * time.Millisecond}
	// strategy := shared.IncrementalTimeoutStrategy{InitialTimeout: 200 * time.Millisecond, Increment: 100 * time.Millisecond}
	strategy := shared.ExponentialBackoffStrategy{InitialBackoff: 200 * time.Millisecond}

	err := shared.Retry(exampleFunction, 5, strategy)
	if err != nil {
		if errors.Is(err, shared.ErrMaxRetriesReached) {
			fmt.Println("Max retries reached. Giving up.")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
		return
	}

	fmt.Println("Success!")
}
