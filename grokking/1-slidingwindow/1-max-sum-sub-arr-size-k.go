package grokking

import "math"

func maxSumSubArr(arr []int, k int) int {
	// maxSumSubArr([]int{2, 1, 5, 1, 3, 2}, 3)

	left := 0
	sum := 0
	maxSum := 0
	for right := 0; right < len(arr); right++ {
		sum += arr[right]

		if right-left+1 == k {
			maxSum = int(math.Max(float64(maxSum), float64(sum)))

			sum -= arr[left]
			left++
		}
	}

	return maxSum
}
