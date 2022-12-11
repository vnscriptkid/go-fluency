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
	// bottomUp([]int{2, 3, 1, 4}, []int{4, 5, 3, 7}, 5)

	//       0  1  2  3  4  5
	//(2,4)  0  0  4  4  4  4
	//(3,5)  0  0  4  5  5  9
	//(1,3)  0  3  4  7  8  9
	//(4,7)  0  3  4  7  8  10

	// Init matrix
	rows := len(weights)
	cols := capacity + 1

	m := make([][]int, rows)

	for i := range m {
		m[i] = make([]int, cols)
	}

	// Init first row
	for i := 1; i <= capacity; i++ {
		if i >= weights[0] {
			m[0][i] = profits[0]
		}
	}

	// Fill matrix
	for i := 1; i < rows; i++ {
		for cap := 1; cap <= capacity; cap++ {
			// First choice: Skip cur item
			p1 := m[i-1][cap]

			p2 := 0

			if cap >= weights[i] {
				p2 = profits[i] + m[i-1][cap-weights[i]]
			}

			m[i][cap] = int(math.Max(float64(p1), float64(p2)))
		}
	}

	return m[rows-1][cols-1]
}
