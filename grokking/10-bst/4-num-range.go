package grokking

func rangeBst(arr []int, key int, leftMost bool) int {
	left := 0
	right := len(arr) - 1

	result := -1

	for left <= right {
		middle := left + (right-left)/2

		if key == arr[middle] {
			result = middle

			if leftMost {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else if key > arr[middle] {
			left = middle + 1
		} else {
			// key < arr[middle]
			right = middle - 1
		}
	}

	return result
}

func numberRange(arr []int, k int) [2]int {

	result := [2]int{-1, -1}
	// 4, 6, 6, 6, 9 | k = 6

	result[0] = rangeBst(arr, k, true)
	result[1] = rangeBst(arr, k, false)

	return result
}
