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
