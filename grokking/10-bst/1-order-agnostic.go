package grokking

func bst(arr []int, key int, asc bool) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := left + (right-left)/2

		if arr[middle] == key {
			return middle
		}

		if asc == true {
			if key > arr[middle] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if key > arr[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}

	return -1
}

func orderAgnosticBst(arr []int, key int) int {
	if arr[0] <= arr[len(arr)-1] {
		return bst(arr, key, true)
	}

	return bst(arr, key, false)
}
