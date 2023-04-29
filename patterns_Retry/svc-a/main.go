package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	lib "github.com/vnscriptkid/go-fluency/patterns_Retry/shared/lib"
)

func requestFunction() (*http.Response, error) {
	resp, err := http.Get("https://example.com")
	if err != nil {
		return nil, err
	}

	// Check status code before returning
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Choose your desired strategy
	// strategy := lib.FixedTimeoutStrategy{Timeout: 200 * time.Millisecond}
	// strategy := lib.IncrementalTimeoutStrategy{InitialTimeout: 200 * time.Millisecond, Increment: 100 * time.Millisecond}
	strategy := lib.ExponentialBackoffStrategy{InitialBackoff: 200 * time.Millisecond}

	resp, err := lib.RetryWithTimeout(requestFunction, 5*time.Second, strategy)
	if err != nil {
		if errors.Is(err, lib.ErrTimeoutReached) {
			fmt.Println("Timeout reached. Giving up.")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Success! Response body:", string(body))
}
