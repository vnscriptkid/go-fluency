package grokking

func findDuplicateNum(arr []int) int {
	// 1  2  3  4  5
	// 1, 2, 3, 4, 4
	//             ^

	i := 0

	for i < len(arr) {
		curNum := arr[i]
		correctIdx := curNum - 1

		if i != correctIdx && arr[correctIdx] != correctIdx+1 {
			arr[i], arr[correctIdx] = arr[correctIdx], arr[i]
		} else {
			i++
		}
	}

	return arr[len(arr)-1]
}
