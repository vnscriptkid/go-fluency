package grokking

import "strconv"

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
