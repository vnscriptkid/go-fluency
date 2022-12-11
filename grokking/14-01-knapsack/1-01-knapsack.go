package grokking

import (
	"math"
	"strconv"
)

func topDownRecursive(weights []int, profits []int, leftCapacity int, idx int, cache map[string]int) int {
	cKey := strconv.Itoa(leftCapacity) + "," + strconv.Itoa(idx)

	_, ok := cache[cKey]

	if ok == false {
		// one fn call, consider 1 item
		if idx == len(weights) {
			return 0
		}

		// 1nd choice: skip it
		r1 := topDownRecursive(weights, profits, leftCapacity, idx+1, cache)

		// 2st choice: take it (if has enough capacity)
		r2 := 0
		if leftCapacity >= weights[idx] {
			r2 = profits[idx] + topDownRecursive(weights, profits, leftCapacity-weights[idx], idx+1, cache)
		}

		cache[cKey] = int(math.Max(float64(r1), float64(r2)))
	}

	return cache[cKey]
}

func topDown(weights []int, profits []int, capacity int) int {
	return topDownRecursive(weights, profits, capacity, 0, map[string]int{})
}

func bottomUp(weights []int, profits []int, capacity int) int {
	return 0
}
