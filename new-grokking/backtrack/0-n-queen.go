package backtrack

const (
	FREE    = 0
	BLOCKED = 1
)

func setRow(board [][]int, x int, value int) {
	n := len(board)

	for j := 0; j < n; j++ {
		board[x][j] = value
	}
}

func setCol(board [][]int, y int, value int) {
	n := len(board)

	for i := 0; i < n; i++ {
		board[i][y] = value
	}
}

func isValidCell(board [][]int, x int, y int) bool {
	n := len(board)

	if x < 0 || x >= n || y < 0 || y >= n {
		return false
	}

	return true
}

func setDiagonal(board [][]int, x int, y int, value int) {
	// go top right
	curX, curY := x, y
	for {
		curX--
		curY++
		if !isValidCell(board, curX, curY) {
			break
		}
		board[curX][curY] = value
	}

	// go top left
	curX, curY = x, y
	for {
		curX--
		curY--
		if !isValidCell(board, curX, curY) {
			break
		}
		board[curX][curY] = value
	}

	// go bottom left
	curX, curY = x, y
	for {
		curX++
		curY--
		if !isValidCell(board, curX, curY) {
			break
		}
		board[curX][curY] = value
	}

	// go bottom right
	curX, curY = x, y
	for {
		curX++
		curY++
		if !isValidCell(board, curX, curY) {
			break
		}
		board[curX][curY] = value
	}

}

func placeQueen(board [][]int, x int, y int, queensLeft int) bool {
	if queensLeft == 0 {
		return true
	}

	if x >= len(board) {
		return false
	}

	if y >= len(board) {
		return placeQueen(board, x+1, 0, queensLeft)
	}

	if board[x][y] == BLOCKED {
		return placeQueen(board, x, y+1, queensLeft)
	}

	// If control comes here
	// It means cur cell (x,y) is valid cell
	// And it's FREE still

	// block cur row x
	setRow(board, x, BLOCKED)

	// block cur col y
	setCol(board, y, BLOCKED)

	// block diagonal line
	setDiagonal(board, x, y, BLOCKED)

	// place queen at next free cell
	if placeQueen(board, x+1, 0, queensLeft-1) {
		return true
	}

	// make cell free again
	board[x][y] = FREE

	setRow(board, x, FREE)

	setCol(board, y, FREE)

	setDiagonal(board, x, y, FREE)

	return false
}

func initBoard(n int) [][]int {
	board := make([][]int, n)

	for i := range board {
		board[i] = make([]int, n)
	}

	return board
}

func nQueen(n int) int {
	count := 0
	board := initBoard(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if placeQueen(board, i, j, n) {
				count++
			}
		}
	}

	return count
}
