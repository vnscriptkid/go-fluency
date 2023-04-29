package lib

import (
	"errors"
	"math/rand"
	"net/http"
	"time"
)

type RetryableFunc func() (*http.Response, error)

var ErrMaxRetriesReached = errors.New("max retries reached")
var ErrTimeoutReached = errors.New("timeout reached")

type BackoffStrategy interface {
	NextDuration(attempt int) time.Duration
}

// //// FixedTimeoutStrategy //////
type FixedTimeoutStrategy struct {
	Timeout time.Duration
	Jitter  float64
}

func (f FixedTimeoutStrategy) NextDuration(attempt int) time.Duration {
	jitter := time.Duration(f.Jitter * float64(f.Timeout) * (rand.Float64() - 0.5) * 2)
	return f.Timeout + jitter
}

//////////////////////////////////

// //// IncrementalTimeoutStrategy //////
type IncrementalTimeoutStrategy struct {
	InitialTimeout time.Duration
	Increment      time.Duration
}

func (i IncrementalTimeoutStrategy) NextDuration(attempt int) time.Duration {
	return i.InitialTimeout + time.Duration(attempt)*i.Increment
}

///////////////////////////////////////

// ////// ExponentialBackoffStrategy ////////
type ExponentialBackoffStrategy struct {
	InitialBackoff time.Duration
}

func (e ExponentialBackoffStrategy) NextDuration(attempt int) time.Duration {
	backoff := e.InitialBackoff * time.Duration(1<<attempt)
	return backoff + time.Duration(rand.Intn(100))*time.Millisecond
}

///////////////////////////////////////////

// Num of retry attempts
func Retry(retryableFunc RetryableFunc, maxRetries int, strategy BackoffStrategy) (*http.Response, error) {
	for i := 0; i < maxRetries; i++ {
		resp, err := retryableFunc()

		// Success
		if err == nil {
			return resp, nil
		}

		// If it's the last attempt, break then return max retries error
		if i == maxRetries-1 {
			break
		}

		// If it's a server error, wait for next retry
		if resp != nil && resp.StatusCode >= 500 {
			sleepDuration := strategy.NextDuration(i)
			time.Sleep(sleepDuration)
		} else {
			// If it's a client error, don't wait for next retry
			return resp, err
		}
	}

	return nil, ErrMaxRetriesReached
}

// Timeboxing
func RetryWithTimeout(retryableFunc RetryableFunc, timeout time.Duration, strategy BackoffStrategy) (resp *http.Response, err error) {
	startTime := time.Now()

	for attempt := 0; ; attempt++ {
		resp, err = retryableFunc()

		if err == nil {
			return resp, nil
		}

		// Measure duration since start time til next retry, is it more than timeout?
		if time.Since(startTime)+strategy.NextDuration(attempt) >= timeout {
			// If yes, break then return timeout error
			break
		}

		// If it's a server error, wait for next retry
		if resp != nil && resp.StatusCode >= 500 {
			sleepDuration := strategy.NextDuration(attempt)
			time.Sleep(sleepDuration)
		} else {
			// If it's a client error, don't wait for next retry
			return resp, err
		}
	}

	return nil, ErrTimeoutReached
}
