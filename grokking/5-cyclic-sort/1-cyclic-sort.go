package grokking

import "fmt"

func cyclicSort(arr []int) {
	// [3, 1, 5, 4, 2]
	i := 0

	for i < len(arr) {
		correctIdx := arr[i] - 1

		if i == correctIdx {
			i++
			continue
		}

		arr[i], arr[correctIdx] = arr[correctIdx], arr[i]
	}
}

func main() {
	// Input: [3, 1, 5, 4, 2]
	// Output: [1, 2, 3, 4, 5]
	arr := []int{3, 1, 5, 4, 2}
	cyclicSort(arr)
	fmt.Println(arr)

	// 	Input: [2, 6, 4, 3, 1, 5]
	// Output: [1, 2, 3, 4, 5, 6]

	arr = []int{2, 6, 4, 3, 1, 5}
	cyclicSort(arr)
	fmt.Println(arr)
}
