package TwoPointers

func isPalindrome(inputString string) bool {
	// Write your code here
	// Tip: You may use the code template provided
	// in the two_pointers.go file
	// your code will replace this placeholder return statement
	left := 0
	right := len(inputString) - 1

	for left <= right {
		if inputString[left] != inputString[right] {
			return false
		}

		left++
		right--
	}

	return true
}
