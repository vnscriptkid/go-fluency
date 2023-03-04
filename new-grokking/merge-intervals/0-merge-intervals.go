package mergeintervals

import (
	"math"
	"strconv"
)

func mergeIntervals(v []Interval) []Interval {
	// [[1, 5], [3, 7], [4, 6]]
	//    ^
	result := make([]Interval, 0)

	prev := v[0]

	for i := 1; i < len(v); i++ {
		cur := v[i]

		// [   ]
		//       {  }

		if cur.start <= prev.end {
			// merge
			prev = *cur.Merge(&prev)
		} else {
			result = append(result, prev)
		}
	}

	result = append(result, prev)

	return result
}

// Template for interval class

type Interval struct {
	start  int
	end    int
	closed bool
}

func (i *Interval) IntervalInit(start int, end int) {
	i.start = start
	i.end = end

	// By default the interval is closed
	i.closed = true
}

func (i *Interval) Merge(o *Interval) *Interval {
	// Assuming 2 overlaps
	return &Interval{
		start: int(math.Min(float64(i.start), float64(o.start))),
		end:   int(math.Max(float64(i.end), float64(o.end))),
	}
}

func (i *Interval) setClosed(closed bool) {
	i.closed = closed
}

func (i *Interval) str() string {
	out := ""
	if i.closed {
		out = "[" + strconv.Itoa(i.start) + ", " + strconv.Itoa(i.end) + "]"
	} else {
		out = "(" + strconv.Itoa(i.start) + ", " + strconv.Itoa(i.end) + ")"
	}
	return out
}
