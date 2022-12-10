package grokking

func subset(arr []int) [][]int {
	result := [][]int{}

	// [1, 5, 3]

	// [ [], [1], [3], [1,3], [5], [1,5], [3,5], [1,3,5] ]
	//                   ^

	result = append(result, []int{})

	for _, num := range arr {
		resultLen := len(result)

		for i := 0; i < resultLen; i++ {
			clone := result[i][:]
			clone = append(clone, num)
			result = append(result, clone)
		}
	}

	return result
}
