package grokking

import (
	"fmt"
	"sort"
)

func tripletSumToZero(arr []int) [][]int {
	// sort array in asc order
	sort.Ints(arr)

	fmt.Println(arr)

	// -3 -2 -1 0 1 1 2
	// @   ^          $

	// fix left most number
	result := [][]int{}

	for firstIdx := 0; firstIdx < len(arr)-2; firstIdx++ {
		secondIdx := firstIdx + 1
		thirdIdx := len(arr) - 1

		for secondIdx < thirdIdx {
			curSum := arr[secondIdx] + arr[thirdIdx]

			if curSum == -arr[firstIdx] {
				result = append(result, []int{arr[firstIdx], arr[secondIdx], arr[thirdIdx]})

				secondIdx++
				thirdIdx--

				for arr[secondIdx] == arr[secondIdx-1] {
					secondIdx++
				}
				for arr[thirdIdx] == arr[thirdIdx+1] {
					thirdIdx--
				}
			} else if curSum < -arr[firstIdx] {
				secondIdx++
			} else {
				thirdIdx--
			}
		}
	}

	return result
}
