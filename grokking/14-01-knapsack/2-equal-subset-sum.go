package grokking

import (
	"strconv"
)

func _equalSubsetSumTopDown(arr []int, curIdx int, target int, cache map[string]bool) bool {
	key := strconv.Itoa(curIdx) + "," + strconv.Itoa(target)

	if _, ok := cache[key]; !ok {
		if target == 0 {
			return true
		}

		if curIdx == len(arr) {
			return false
		}

		// Skip or Take
		cache[key] = (_equalSubsetSumTopDown(arr, curIdx+1, target-arr[curIdx], cache) || _equalSubsetSumTopDown(arr, curIdx+1, target, cache))
	}

	return cache[key]
}

func equalSubsetSumTopDown(arr []int) bool {
	sum := 0
	for _, x := range arr {
		sum += x
	}

	if sum%2 == 1 {
		return false
	}

	half := sum / 2

	return _equalSubsetSumTopDown(arr, 0, half, map[string]bool{})
}

func equalSubsetSumBottomUp(arr []int) bool {
	sum := 0
	for _, x := range arr {
		sum += x
	}

	if sum%2 == 1 {
		return false
	}

	half := sum / 2

	dp := make([][]bool, len(arr))
	for i := range dp {
		dp[i] = make([]bool, half+1)
	}

	//    0  1  2  3  4  5
	// 1  T  T  f  f  f  f
	// 2  T  T  T  T  f  f
	// 3  T  T  T  T  T  T
	// 4  T  T  T  T  T  T

	// Init first col
	for i := range dp {
		dp[i][0] = true
	}

	// Init first row
	if arr[0] > 0 && arr[0] <= half {
		dp[0][arr[0]] = true
	}

	for i := 1; i < len(arr); i++ {
		for j := 1; j <= half; j++ {
			if dp[i-1][j] == true {
				dp[i][j] = true
				continue
			}

			if j >= arr[i] {
				dp[i][j] = dp[i-1][j-arr[i]]
			}
		}
	}

	return dp[len(dp)-1][half]
}
