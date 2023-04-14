package bst

func binarySearchRotated(nums []int, target int) int {
	// 6, 7,| 1, 2, 3, 4, 5, target:7

	// left, middle, right

	// arr[middle] <= arr[right] && target <= arr[right] && arr[middle] < target
	// => [ middle...right ] => left = middle + 1

	// arr[left] <= target <= arr[middle]
	// => [left...middle] => right = middle - 1

	// arr[middle] < arr[left] => [middle...left]
	// arr[middle] > arr[right] => [middle...right]
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + (right-left)/2

		if target == nums[middle] {
			return middle
		}

		if nums[left] <= target && target < nums[middle] {
			// [left...middle)
			right = middle - 1
			continue
		}

		if nums[middle] < target && target <= nums[right] {
			// (middle...right]
			left = middle + 1
			continue
		}

		if nums[left] > nums[middle] {
			right = middle - 1
			continue
		}

		if nums[middle] > nums[right] {
			left = middle + 1
			continue
		}

		return -1
	}

	// write your code here
	// your code will replace this placeholder return statement
	return -1
}
