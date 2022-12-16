package grokking

func ceilingOfNum(arr []int, key int) int {
	// 1, 3, 8, 10, 15

	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := left + (right-left)/2

		if arr[middle] == key {
			return middle
		}

		if key > arr[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	// Now left > right

	if left == len(arr) {
		return -1
	}

	return left
}
