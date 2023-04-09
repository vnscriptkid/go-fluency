package Track

func permutePalindrome(st string) bool {
	count := map[rune]int{}

	for _, c := range st {
		if _, ok := count[c]; !ok {
			count[c] = 0
		}

		count[c] += 1
	}

	oddCharsCount := 0

	for _, v := range count {
		if v%2 == 1 {
			oddCharsCount++
		}
	}

	// Write your code here
	return oddCharsCount <= 1
}
