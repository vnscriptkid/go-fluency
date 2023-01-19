package grokking

func longestSubstrKDistinctChars(str string, k int) int {
	// a  r  a  a  c  i  |  k = 2
	//       ^     $

	// { a: 2, c: 1 }

	// distinctChars: 3
	// longest: 4 (windowSize)

	start := 0
	charMap := map[rune]int{}
	longest := 1

	for end := 0; end < len(str); end++ {
		// expand window
		curChar := (rune)(str[end])
		if _, ok := charMap[curChar]; !ok {
			charMap[curChar] = 0
		}
		charMap[curChar] += 1

		// the expansion can make string invalid (distinct chars > k)
		// loop to bring window to valid state
		for len(charMap) > k {
			// shrink down window from the left
			leftMostChar := (rune)(str[start])

			charMap[leftMostChar] -= 1

			if charMap[leftMostChar] == 0 {
				// remove this from the map
				delete(charMap, leftMostChar)
			}
			start += 1
		}

		// update longest if possible
		curWindowSize := end - start + 1

		if curWindowSize > longest {
			longest = curWindowSize
		}
	}

	return longest
}
