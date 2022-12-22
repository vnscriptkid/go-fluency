package grokking

import "math"

func smallestSubArrWithSum(arr []int, sum int) int {
	// [2, 1, 5, 2, 3, 2] sum: 7
	//           ^     $

	// curSum 7

	// curMin 2

	left := 0
	curSum := 0
	curMin := math.Inf(1)

	for right := 0; right < len(arr); right++ {
		curSum += arr[right]

		for curSum >= sum {
			windowSize := right - left + 1

			curMin = math.Min(curMin, float64(windowSize))

			curSum -= arr[left]
			left++
		}
	}

	return int(curMin)
}
