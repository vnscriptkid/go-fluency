package grokking

import (
	"strconv"
)

func recur(subset []int, curIdx int, targetS int, cache map[string]bool) bool {
	cacheKey := strconv.Itoa(curIdx) + "," + strconv.Itoa(targetS)

	if _, ok := cache[cacheKey]; !ok {
		if targetS == 0 {
			return true
		}

		if curIdx == len(subset) {
			return false
		}

		cache[cacheKey] = recur(subset, curIdx+1, targetS-subset[curIdx], cache) ||
			recur(subset, curIdx+1, targetS, cache)
	}

	return cache[cacheKey]
}

func subsetSumTopdown(subset []int, s int) bool {

	return recur(subset, 0, s, map[string]bool{})
}

func subsetSumBottomUp(subset []int, s int) bool {
	//    0  1  2  3  4  5  6
	// 1  x  x  _  _  _  _  _
	// 2  x  x  x  x  _  _  _
	// 3  x  x  x  x  x  x  x
	// 7  x

	dp := make([][]bool, len(subset))

	for i := range dp {
		dp[i] = make([]bool, s+1)
	}

	for i := 0; i < len(subset); i++ {
		for j := 0; j <= s; j++ {
			if j == 0 || j == subset[i] {
				dp[i][j] = true
				continue
			}

			dp[i][j] = false
		}
	}

	for i := 1; i < len(subset); i++ {
		for j := 1; j <= s; j++ {
			if dp[i-1][j] == true {
				dp[i][j] = true
				continue
			}

			if j >= subset[i] {
				dp[i][j] = dp[i-1][j-subset[i]]
			}
		}
	}

	return dp[len(subset)-1][s]
}
