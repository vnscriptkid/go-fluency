package main

import (
	"fmt"
	"math"
)

// [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
//  ^
//              $

// [2.2, 2.8, 2.4, 3.6, 2.8]

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func avgSubArrSizeK(arr []int, k int) []float64 {

	left := 0
	avg := float64(0)
	result := []float64{}

	for right := range arr {
		avg += float64(arr[right]) / float64(k)

		if right-left+1 == k {
			result = append(result, roundFloat(avg, 1))
			avg -= float64(arr[left]) / float64(k)
			left++
		}
	}

	return result
}

func main() {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5

	fmt.Println(avgSubArrSizeK(arr, k))
}
