package grokking

func makeSquares(arr []int) []int {
	//-2, -1, 0, 2, 3
	//        ^$

	// 0 vs 0
	//     0   1  4   4  9

	left := 0
	right := len(arr) - 1
	result := []int{}

	for left <= right {
		squareLeft := arr[left] * arr[left]
		squareRight := arr[right] * arr[right]

		if squareLeft > squareRight {
			result = append([]int{squareLeft}, result...)
			left++
		} else {
			result = append([]int{squareRight}, result...)
			right--
		}
	}

	return result
}
