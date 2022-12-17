package grokking

func findSingleNumbers(arr []int) (int, int) {
	axb := 0

	for _, x := range arr {
		axb ^= x
	}

	// find the diff bit
	setBit := 1

	for setBit&axb == 0 {
		setBit = setBit << 1
	}

	// partition
	a, b := 0, 0
	for _, x := range arr {
		if setBit&x == 0 {
			a ^= x
		} else {
			b ^= x
		}
	}

	return a, b
}
