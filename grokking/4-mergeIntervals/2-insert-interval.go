package grokking

import "math"

type MyInterval struct {
	Start int
	End   int
}

func insertInterval(intervals []MyInterval, newInterval MyInterval) []MyInterval {
	// 1  2  3  4  5  6  7  8  9 10 11 12
	// [     ]
	//             [     ]
	//                      [           ]
	//               ^
	//          {     }

	// 3 periods
	// in match
	// post match
	r := []MyInterval{}

	// pre match
	i := 0

	for intervals[i].End < newInterval.Start {
		r = append(r, intervals[i])
		i++
	}

	merged := newInterval

	for i < len(intervals) && merged.End >= intervals[i].Start {
		merged = MyInterval{
			Start: int(math.Min(
				float64(merged.Start),
				float64(intervals[i].Start),
			)),
			End: int(math.Max(
				float64(merged.End),
				float64(intervals[i].End),
			)),
		}
		i++
	}

	r = append(r, merged)

	// post match
	for i < len(intervals) {
		r = append(r, intervals[i])
		i++
	}

	return r
}
