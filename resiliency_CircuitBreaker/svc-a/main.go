package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type CircuitBreaker struct {
	State        int
	Counter      int
	OpenTime     time.Time
	ResetTimeout time.Duration
	MaxFailures  int
}

const (
	StateClosed = iota
	StateOpen
	StateHalfOpen
)

func (cb *CircuitBreaker) ExecuteRequest(req *http.Request) (*http.Response, error) {

	if cb.State == StateOpen {
		if time.Since(cb.OpenTime) >= cb.ResetTimeout {
			cb.State = StateHalfOpen
		} else {
			return nil, errors.New("circuit breaker is open")
		}
	}
	client := &http.Client{Timeout: time.Second}
	resp, err := client.Do(req)
	if err != nil {
		cb.Counter++
		if cb.Counter >= cb.MaxFailures {
			cb.State = StateOpen
			cb.OpenTime = time.Now()
		}
		return nil, err
	}
	cb.Counter = 0
	cb.State = StateClosed
	return resp, nil
}

func main() {
	fmt.Println("svc-a started")

	// Create a new circuit breaker for ServiceA with the desired settings
	cbB := CircuitBreaker{
		State:        StateClosed,
		Counter:      0,
		OpenTime:     time.Time{},
		ResetTimeout: time.Duration(10) * time.Second,
		MaxFailures:  3,
	}

	// Call ServiceB through the circuit breaker
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequest("GET", "https://servicea.example.com/api", nil)
		resp, err := cbB.ExecuteRequest(req)
		if err != nil {
			fmt.Printf("Error calling ServiceA: %s\n", err)
		} else {
			fmt.Printf("Response status from ServiceA: %s\n", resp.Status)
		}
	}

	// Create a new circuit breaker for ServiceC with the desired settings
	cbC := CircuitBreaker{
		State:        StateClosed,
		Counter:      0,
		OpenTime:     time.Time{},
		ResetTimeout: time.Duration(10) * time.Second,
		MaxFailures:  3,
	}

	// Call ServiceC through the circuit breaker
	for i := 0; i < 4; i++ {
		req, _ := http.NewRequest("GET", "https://serviceb.example.com/api", nil)
		resp, err := cbC.ExecuteRequest(req)
		if err != nil {
			fmt.Printf("Error calling ServiceB: %s\n", err)
		} else {
			fmt.Printf("Response status from ServiceB: %s\n", resp.Status)
		}
	}
}
