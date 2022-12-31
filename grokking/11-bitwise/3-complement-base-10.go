package grokking

import "math"

func bitwiseComplement(n int) int {
	// n ^ complement = allBitSet

	// n ^ n ^ complement = allBitSet ^ n

	// complement = allBitSet ^ n

	bitsCount := 0

	nValue := n

	for nValue > 0 {
		bitsCount++
		nValue = nValue >> 1 // right shift one bit
	}

	allBitSet := int(math.Pow(float64(2), float64(bitsCount)) - 1)

	return allBitSet ^ n
}
