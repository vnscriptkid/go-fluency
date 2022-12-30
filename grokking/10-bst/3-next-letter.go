package grokking

func nextLetter(letters []rune, key rune) rune {
	// 'a', 'c', 'f', 'h', key: 'f

	left := 0
	right := len(letters) - 1

	for left <= right {
		middle := left + (right-left)/2

		if key == letters[middle] {
			return letters[(middle+1)%len(letters)]
		}

		if key > letters[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	// left > right now
	return letters[left%len(letters)]
}
