package grokking

func singleNumber(arr []int) int {
	r := 0

	for _, num := range arr {
		r = r ^ num
	}

	return r
}
