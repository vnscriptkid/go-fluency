package cyclicSort

func findMissingNumber(nums []int) int {
	// 0  1  2  3  4
	// 0, 1, 2, 4
	//            ^

	// First round: bring number to it's correct idx
	// number x at index x

	i := 0
	for i < len(nums) {
		num := nums[i]

		// If num is largest num OR
		// num is already at correct place
		if num == i || num == len(nums) {
			// Then Move on
			i++
		} else {
			// Cur num is not at correct place
			// Swap position with number at it's correct idx
			nums[i], nums[num] = nums[num], nums[i]
		}
	}

	// Second round
	for j, num := range nums {
		if j != num {
			return j
		}
	}

	return len(nums)
}
