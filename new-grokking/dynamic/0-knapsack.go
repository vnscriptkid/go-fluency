package dynamic

import "math"

func findMaxKnapsackProfit(capacity int, weights []int, values []int) int {
	// Init dp
	dp := make([][]int, len(weights))

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Init first row
	for curCap := range dp[0] {
		if curCap >= weights[0] {
			dp[0][curCap] = values[0]
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j <= capacity; j++ {
			// Choice 1: Skip this item
			dp[i][j] = dp[i-1][j]

			// Choice 2: Take this item
			if j >= weights[i] {
				profitIfTake := values[i] + dp[i-1][j-weights[i]]

				dp[i][j] = int(math.Max(
					float64(dp[i][j]), float64(profitIfTake),
				))
			}
		}
	}

	return dp[len(dp)-1][capacity]
}
