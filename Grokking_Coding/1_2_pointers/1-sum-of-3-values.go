package TwoPointers

import "sort"

func findSumOfThree(nums []int, target int) bool {
	// Sort the array in ascending order
	sort.Ints(nums)

	// Loop through the array, fix each element as the first element of the triplet

	for i := 0; i < len(nums)-2; i++ {
		// Find the other two elements using two pointers
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return true
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return false
}
