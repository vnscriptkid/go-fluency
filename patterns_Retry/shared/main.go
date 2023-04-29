package main

import (
	"errors"
	"math/rand"
	"time"
)

type RetryableFunc func() error

var ErrMaxRetriesReached = errors.New("max retries reached")

type BackoffStrategy interface {
	NextDuration(attempt int) time.Duration
}

type FixedTimeoutStrategy struct {
	Timeout time.Duration
}

func (f FixedTimeoutStrategy) NextDuration(attempt int) time.Duration {
	return f.Timeout
}

type IncrementalTimeoutStrategy struct {
	InitialTimeout time.Duration
	Increment      time.Duration
}

func (i IncrementalTimeoutStrategy) NextDuration(attempt int) time.Duration {
	return i.InitialTimeout + time.Duration(attempt)*i.Increment
}

type ExponentialBackoffStrategy struct {
	InitialBackoff time.Duration
}

func (e ExponentialBackoffStrategy) NextDuration(attempt int) time.Duration {
	backoff := e.InitialBackoff * time.Duration(1<<attempt)
	return backoff + time.Duration(rand.Intn(100))*time.Millisecond
}

func Retry(retryableFunc RetryableFunc, maxRetries int, strategy BackoffStrategy) error {
	var err error

	for i := 0; i < maxRetries; i++ {
		err = retryableFunc()
		if err == nil {
			return nil
		}

		sleepDuration := strategy.NextDuration(i)
		time.Sleep(sleepDuration)
	}

	return ErrMaxRetriesReached
}
