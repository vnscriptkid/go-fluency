package grokking

import "fmt"

// Given an array of sorted numbers and a target sum,
// find a pair in the array whose sum is equal to the given target.

// 1  2  3  4  6
//    ^
//          $

// sum: 6 => (1,3)
// target: 6

func twoSum(arr []int, targetSum int) (int, int) {

	left := 0
	right := len(arr) - 1
	var curSum int

	for left < right {
		curSum = arr[left] + arr[right]

		if curSum == targetSum {
			return left, right
		}

		if curSum < targetSum {
			left++
		} else {
			right--
		}
	}

	return -1, -1
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 6}, 6))
}
