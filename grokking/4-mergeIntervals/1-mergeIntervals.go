package main

import (
	"fmt"
	"math"
	"sort"
)

type Interval struct {
	open int
	end  int
}

// 1 2 3 4 5 6 7 8 9
// [     ]
//    p
//   [     ]
//      c
//             [   ]

func mergeIntervals(intervals []Interval) []Interval {
	// Sort by open time in asc order in-place
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].open < intervals[j].open
	})

	result := []Interval{}

	prev := intervals[0]

	for i, cur := range intervals {
		// Skip i === 0 as cur starts from idx 1
		if i == 0 {
			continue
		}

		// if cur overlaps prev => merge cur with prev
		// reassign merged to prev
		if cur.open <= prev.end {
			prev = Interval{
				// open: math.MinInt(prev.open, cur.open),
				// math.MaxInt(prev.close, cur.close),
				open: int(math.Min(float64(prev.open), float64(cur.open))),
				end:  int(math.Max(float64(prev.end), float64(cur.end))),
			}
		} else {
			// if cur does NOT overlaps prev
			// reassign cur to prev
			result = append(result, prev)
			prev = cur
		}

	}

	result = append(result, prev)

	return result
}

func main() {
	// 	Intervals: [[1,4], [2,5], [7,9]]
	// Output: [[1,5], [7,9]]
	fmt.Println(mergeIntervals([]Interval{
		{1, 4},
		{2, 5},
		{7, 9},
	}))

	// 	Intervals: [[6,7], [2,4], [5,9]]
	// Output: [[2,4], [5,9]]
	fmt.Println(mergeIntervals([]Interval{
		{6, 7},
		{2, 4},
		{5, 9},
	}))

	// Intervals: [[1,4], [2,6], [3,5]]
	// Output: [[1,6]]
	fmt.Println(mergeIntervals([]Interval{
		{1, 4},
		{2, 6},
		{3, 5},
	}))
}
