package TraverseAllVertices

import "log"

func dfs(g [][]int, v []int, visited map[int]bool, c int) []int {
	// Decide to visit vertex c or not
	if seen := visited[c]; seen {
		return v
	}

	// Now visit it
	v = append(v, c)
	visited[c] = true

	// Check through all vertices connected to vertex c
	for i := 0; i < len(g); i++ {
		if g[c][i] == 1 && c != i {
			v = dfs(g, v, visited, i)
		}
	}

	return v
}

func traverse(g [][]int) []int {
	v := []int{}

	// Init visited map
	visited := map[int]bool{}

	n := len(g)

	for i := 0; i < n; i++ {
		visited[i] = false
	}

	return dfs(g, v, visited, 0)
}

// Time: O(V + E)
// Space: O(V) -> height of stack
// - Max is V when one vertex is connected to all others
func dfsV2(g [][]int, c int, visited []bool) {
	visited[c] = true
	log.Printf("Visited %v", c)

	// Check through all vertices connected to vertex c
	for i := 0; i < len(g); i++ {
		if g[c][i] == 1 && c != i && !visited[i] {
			dfsV2(g, i, visited)
		}
	}
}

func traverseV2(g [][]int) {
	// Init visited arr
	n := len(g)
	visited := make([]bool, n)

	dfsV2(g, 0, visited)
}
