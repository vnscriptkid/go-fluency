package topoSort

import (
	"container/list"
)

func findOrder(dependencies [][]rune) []rune {
	sortedOrder := make([]rune, 0)
	graph := map[rune][]rune{}
	incommings := map[rune]int{}
	queue := list.New()

	for _, pair := range dependencies {
		to, from := pair[0], pair[1]

		if _, ok := graph[from]; !ok {
			graph[from] = make([]rune, 0)
		}
		if _, ok := graph[to]; !ok {
			graph[to] = make([]rune, 0)
		}

		graph[from] = append(graph[from], to)

		if _, ok := incommings[from]; !ok {
			incommings[from] = 0
		}
		if _, ok := incommings[to]; !ok {
			incommings[to] = 0
		}

		incommings[to] = incommings[to] + 1
	}

	for to, count := range incommings {
		if count == 0 {
			queue.PushBack(to)
		}
	}

	for queue.Len() > 0 {
		out := queue.Remove(queue.Front()).(rune)

		sortedOrder = append(sortedOrder, out)

		for _, to := range graph[out] {
			incommings[to] -= 1

			if incommings[to] == 0 {
				queue.PushBack(to)
			}
		}
	}

	// write your code here
	if len(sortedOrder) == len(graph) {
		return sortedOrder
	}

	return make([]rune, 0)
}
