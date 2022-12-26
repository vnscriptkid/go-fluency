package grokking

func findPermutations(nums []int) [][]int {
	// 1 2 3
	// ^

	// 1    1 .   _
	//      /\ .
	// 2 .  2 _

	//              [1]

	//   [3 1]                [1 3]

	// [531] [351] [315] . [513] . [153] . [135]

	// [1 _ 2] , 3, inject at 1
	// [:1] [3] [1:]

	result := [][]int{{nums[0]}}

	for i := 1; i < len(nums); i++ {
		temp := [][]int{}
		num := nums[i]

		// for j := 0; j <= len(result); j++ {
		// 	newPerm :=
		// }
		levelSize := len(result)

		for j := 0; j < levelSize; j++ {
			arr := result[j]

			for k := 0; k <= len(arr); k++ {
				newArr := append(arr[:k], append([]int{num}, arr[k:]...)...)

				temp = append(temp, newArr)
			}
		}

		result = temp
	}

	return result
}
