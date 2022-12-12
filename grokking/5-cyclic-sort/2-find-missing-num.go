package grokking

func findMissingNum(a []int) int {
	// 0, 1, 4, 3
	//       ^

	i := 0

	for i < len(a) {
		correctIdx := a[i]

		if correctIdx < len(a) && i != correctIdx {
			a[i], a[correctIdx] = a[correctIdx], a[i]
		} else {
			i++
		}
	}

	for i, x := range a {
		if i != x {
			return i
		}
	}

	return len(a)
}
