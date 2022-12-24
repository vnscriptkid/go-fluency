package grokking

func findAllMissingNums(nums []int) []int {
	// 0  1  2  3  4  5  6  7
	// 1, 2, 3, 1, 5, 3, 2, 8
	//                   ^

	// 4, 6, 7

	// first round: put number at it's correct pos
	i := 0

	for i < len(nums) {
		correctIdx := nums[i] - 1

		if i != correctIdx && nums[correctIdx] != correctIdx+1 {
			nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
		} else {
			i++
		}
	}

	// second round: find missing
	result := []int{}

	for i, num := range nums {
		if num != i+1 {
			result = append(result, i+1)
		}
	}

	return result
}
