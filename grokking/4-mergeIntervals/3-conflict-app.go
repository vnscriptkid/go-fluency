package grokking

import "sort"

func canAttendAllApp(app [][2]int) bool {
	sort.Slice(app, func(i, j int) bool {
		return app[i][0] < app[j][0]
	})

	prev := app[0]

	for i := 1; i < len(app); i++ {
		cur := app[i]

		// [   ]
		//    {  }

		if cur[0] < prev[1] {
			return false
		}

		prev = cur
	}

	return true
}
