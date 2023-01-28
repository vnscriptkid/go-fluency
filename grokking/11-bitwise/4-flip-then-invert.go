package grokking

func flipThenInvert(m [][]int) [][]int {
	for _, row := range m {
		left := 0
		right := len(row) - 1

		for left <= right {
			// invert
			row[left], row[right] = row[right], row[left]

			// flip
			row[left], row[right] = row[left]^1, row[right]^1
			left++
			right--
		}
	}

	return m
}
