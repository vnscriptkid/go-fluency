package ShortestPath

type Graph struct {
	dest map[string][]string
}

func findShortestPath(edges [][2]string, src string, dst string) int {

	g := &Graph{
		dest: map[string][]string{},
	}

	seen := map[string]bool{}

	// Build graph
	for _, pair := range edges {
		x, y := pair[0], pair[1]

		if _, ok := g.dest[x]; !ok {
			g.dest[x] = []string{}
		}
		if _, ok := g.dest[y]; !ok {
			g.dest[y] = []string{}
		}

		seen[x] = false
		seen[y] = false

		g.dest[x] = append(g.dest[x], y)
		g.dest[y] = append(g.dest[y], x)
	}

	// Use queue to traverse in BFS fashion
	// Keep track of current #layer
	// Stop when meet dst
	layer := 1
	queue := []string{src}

	for len(queue) > 0 {
		// Process all items in one layer
		// Remember layer size
		n := len(queue)

		for n > 0 {
			x := queue[0]
			// Pop from head of queue
			queue = queue[1:]

			if x == dst {
				return layer
			}

			// Check connected nodes of x
			if v, ok := g.dest[x]; ok {
				for _, i := range v {
					if !seen[i] {
						queue = append(queue, i)
					}
				}
			}
		}

		layer += 1
	}

	return 0
}
