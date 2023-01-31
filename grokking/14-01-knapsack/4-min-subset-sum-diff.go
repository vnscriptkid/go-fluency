package grokking

import (
	"math"
	"strconv"
)

func minSubsetSumDiffRecur(subset []int, curIdx int, curSum int, globalMin *int, total int, cache map[string]int) int {
	cacheKey := strconv.Itoa(curIdx) + "," + strconv.Itoa(curSum)

	if _, ok := cache[cacheKey]; !ok {
		if curIdx == len(subset) {
			// done with current branch
			other := total - curSum

			diff := math.Abs(float64(curSum - other))

			if int(diff) < *globalMin {
				*globalMin = int(diff)
			}

			return *globalMin
		}

		minSubsetSumDiffRecur(subset, curIdx+1, curSum+subset[curIdx], globalMin, total, map[string]int{})

		minSubsetSumDiffRecur(subset, curIdx+1, curSum, globalMin, total, map[string]int{})

		cache[cacheKey] = *globalMin
	}

	return cache[cacheKey]
}

func minSubsetSumDiff(subset []int) int {
	globalMin := math.MaxInt

	total := 0

	for _, x := range subset {
		total += x
	}

	return minSubsetSumDiffRecur(subset, 0, 0, &globalMin, total, map[string]int{})

}

func minSubsetSumDiffBottomUp(subset []int) int {
	total := 0

	for _, x := range subset {
		total += x
	}

	half := int(math.Ceil(float64(total) / 2))

	var dp [][]bool

	// 1,2,3,9 => 15/2 ~ 8

	//    0  1  2  3  4  5  6  7  8
	// 1  x  x  _  _  _  _  _  _  _
	// 2  x
	// 3  x
	// 9  x

	for i := 0; i < len(subset); i++ {
		dp = make([][]bool, len(subset))
	}

	for i := range dp {
		dp[i] = make([]bool, half+1)
	}

	for i := range dp {
		dp[i][0] = true
	}

	if subset[0] <= half {
		dp[0][subset[0]] = true
	}

	for i := 1; i < len(subset); i++ {
		for j := 1; j <= half; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
				continue
			}

			if j >= subset[i] {
				dp[i][j] = dp[i-1][j-subset[i]]
			}
		}
	}

	for j := half; j >= 0; j-- {
		if dp[len(subset)-1][j] {
			other := total - j
			diff := int(math.Abs(float64(other - j)))

			return diff
		}
	}

	return -1
}
