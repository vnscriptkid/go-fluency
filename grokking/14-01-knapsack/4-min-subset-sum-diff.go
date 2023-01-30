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
