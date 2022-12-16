package grokking

func findSubsetsWDuplicates(a []int) [][]int {
	r := [][]int{}

	r = append(r, []int{})

	// 1, 3, 3
	//       ^

	//      []

	//       []   [1]

	//    []   [1]   [3]  [1,3]

	//   []   [1]   [3]  [1,3]  [3,3]  [1,3,3]

	startIdx := 0

	for idx, x := range a {
		endIdx := len(r) - 1

		for i := startIdx; i <= endIdx; i++ {
			clone := r[i][:]
			clone = append(clone, x)
			r = append(r, clone)
		}

		if idx+1 == len(a) {
			break
		}

		if x != a[idx+1] {
			startIdx = 0
		} else {
			startIdx = endIdx + 1
		}
	}

	return r
}
