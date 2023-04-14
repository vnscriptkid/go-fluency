package backtrack

import "math"

type position struct {
	x int
	y int
}

func solve(curRow int, n int, posList []position, count *int) {
	if curRow == n {
		*count = *count + 1
		return
	}

	for col := 0; col < n; col++ {
		// try to place queen at A[curRow][col]
		// check if queen is under attack by considering every previously placed queens
		// which inside posList
		attacked := false

		for j := 0; j < curRow; j++ {
			pos := posList[j]
			// need checking
			// 1) if under attack vertically by pos
			if pos.y == col {
				attacked = true
				break
			}
			// 2) if under attack diagnolly by pos
			diffX := math.Abs(float64(pos.x - curRow))
			diffY := math.Abs(float64(pos.y - col))
			if diffX == diffY {
				attacked = true
				break
			}
		}

		if attacked {
			continue
		}

		// not attached
		posList[curRow] = position{x: curRow, y: col}
		solve(curRow+1, n, posList, count)
	}
}

func nQueen(n int) int {
	count := 0
	posList := make([]position, n)

	solve(0, n, posList, &count)

	return count
}
