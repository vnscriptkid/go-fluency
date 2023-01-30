package grokking

import "math"

func minSubsetSumDiffRecur(subset []int, curIdx int, curSum int, globalMin *int, total int) {
	if curIdx == len(subset) {
		// done with current branch
		other := total - curSum

		diff := math.Abs(float64(curSum - other))

		if int(diff) < *globalMin {
			*globalMin = int(diff)
		}

		return
	}

	minSubsetSumDiffRecur(subset, curIdx+1, curSum+subset[curIdx], globalMin, total)

	minSubsetSumDiffRecur(subset, curIdx+1, curSum, globalMin, total)
}

func minSubsetSumDiff(subset []int) int {
	globalMin := math.MaxInt

	total := 0

	for _, x := range subset {
		total += x
	}

	minSubsetSumDiffRecur(subset, 0, 0, &globalMin, total)

	return globalMin
}
