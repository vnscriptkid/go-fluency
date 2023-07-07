package PathExists

/*
	Though process:
	Question is yes/no, dont have to check all cases
	Use DFS to go deep and try all possible routes, stop when found
	Use flags to mark visited nodes (visit each nodes once) -> O(V)
	- Build graph:
		- [v] Adjacency list
		- Adjacency matrix
	- Start from src node
		- Visit node (mark visited)
		- Check all connected nodes
		- Go to one of the node and repeat using DFS
		- Stop when found dst node

*/

type Vertex struct {
	key     int
	friends []*Vertex
}

type Graph struct {
	// v *[]Vertex <<<< pointer to slice of vertices
	// v []Vertex <<<< slice of vertices
	v []*Vertex // <<<< slice of vertex pointers
}

// Time: O(V)
// Space: O(V)
func validPath(n int, edges [][]int, src int, dst int) bool {
	v := make([]*Vertex, n)

	g := &Graph{
		v: v,
	}

	for i := range g.v {
		g.v[i] = &Vertex{}
		g.v[i].key = i
		g.v[i].friends = make([]*Vertex, 0)
	}

	for _, pair := range edges {
		x, y := pair[0], pair[1]

		// Add y to x's friends list
		g.v[x].friends = append(g.v[x].friends, g.v[y])
		// Add x to y's friends list
		g.v[y].friends = append(g.v[y].friends, g.v[x])
	}

	visited := make([]bool, n)

	return dfs(g, src, dst, visited)
}

func dfs(g *Graph, cur int, dst int, visited []bool) bool {
	// Visit cur
	visited[cur] = true

	if cur == dst {
		return true
	}

	// Check all cur's friends
	for i := range g.v[cur].friends {
		f := g.v[cur].friends[i].key

		if !visited[f] {
			found := dfs(g, f, dst, visited)

			if found {
				return true
			}
		}
	}

	return false
}
