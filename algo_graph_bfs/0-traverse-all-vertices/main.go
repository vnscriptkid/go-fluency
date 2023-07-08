package TraverseAllVertices

// Time: O(V + E)
// Space: O(V) -> max queue size
func bfs(g [][]int, queue []int, seen []bool, order []int) []int {
	for len(queue) > 0 {
		x := queue[0]
		// Pop from head
		queue = queue[1:]

		// Safety check
		if seen[x] {
			continue
		}

		// Visit x
		seen[x] = true
		order = append(order, x)

		// Check for connected nodes with x
		for i := 0; i < len(g); i++ {
			if i != x && g[i][x] == 1 && !seen[i] {
				queue = append(queue, i)
			}
		}
	}

	return order
}

func traverse(g [][]int) []int {
	order := make([]int, 0)

	queue := make([]int, 0)

	seen := make([]bool, len(g))

	// Start from node 0
	queue = append(queue, 0)

	return bfs(g, queue, seen, order)
}
