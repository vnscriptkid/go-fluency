package grokking

import "container/list"

type Edge struct {
	From int
	To   int
}

func topoSort(vertices int, edges []Edge) []int {
	graph := map[int][]int{}
	incommings := map[int]int{}

	// Init graph
	for v := 0; v < vertices; v++ {
		graph[v] = []int{}
		incommings[v] = 0
	}

	for _, e := range edges {
		graph[e.From] = append(graph[e.From], e.To)
		incommings[e.To] += 1
	}

	// Where to start
	sources := list.List{}
	for v, i := range incommings {
		if i == 0 {
			sources.PushBack(v)
		}
	}

	result := []int{}

	for sources.Len() > 0 {
		v := sources.Remove(sources.Front()).(int)

		result = append(result, v)

		for _, d := range graph[v] {
			incommings[d] -= 1

			if incommings[d] == 0 {
				sources.PushBack(d)
			}
		}
	}

	return result
}
