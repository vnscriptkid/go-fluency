package grokking

import (
	"strings"
	"unicode"
)

func permutationChangingCase(str string) []string {
	// 	Input: "ab7c"
	// Output: "ab7c", "Ab7c", "aB7c", "AB7c", "ab7C", "Ab7C", "aB7C", "AB7C"

	//            [a b 7 c]
	//                   ^

	//   [ab7c   Ab7c  aB7c   AB7c]

	result := []string{}

	result = append(result, str)

	for i, char := range str {
		if !unicode.IsLetter(char) {
			continue
		}

		lenNow := len(result)

		for j := 0; j < lenNow; j++ {
			curStr := result[j]

			result = append(result, curStr[:i]+strings.ToUpper(string(curStr[i]))+curStr[i+1:])
		}
	}

	return result
}
