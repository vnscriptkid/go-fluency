package TraverseAllPaths

func dfs(g [][]int, p []int, a [][]int, now int, dest int, seen []bool) ([][]int, []int) {
	// Visit now
	seen[now] = true
	p = append(p, now)

	if now == dest {
		// Store current path to a
		p2 := make([]int, len(p))
		copy(p2, p)
		a = append(a, p2)
		return a, p
	}

	// Find connected node
	for i := 0; i < len(g); i++ {
		if i != now && g[i][now] == 1 && !seen[i] {
			a, p = dfs(g, p, a, i, dest, seen)

			// After visiting i, reset seen at i
			seen[i] = false
			// Pop i out of current path (final element in path)
			p = p[:len(p)-1]
		}
	}

	return a, p
}

func traverse(g [][]int, f int, t int) [][]int {
	a := [][]int{}
	p := []int{}
	seen := make([]bool, len(g))

	a, _ = dfs(g, p, a, f, t, seen)

	return a
}
