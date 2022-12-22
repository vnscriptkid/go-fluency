package grokking

import "container/list"

func taskScheduling(tasks int, prerequisites [][2]int) bool {
	graph := map[int][]int{}
	incommings := map[int]int{}

	for task := 0; task < tasks; task++ {
		graph[task] = []int{}
		incommings[task] = 0
	}

	for _, pair := range prerequisites {
		from, to := pair[0], pair[1]

		graph[from] = append(graph[from], to)
		incommings[to] += 1
	}

	sources := list.List{}
	ordering := list.List{}

	for toTask, deps := range incommings {
		if deps == 0 {
			sources.PushBack(toTask)
		}
	}

	for sources.Len() > 0 {
		next := sources.Remove(sources.Front()).(int)

		ordering.PushBack(next)

		for _, toTask := range graph[next] {
			incommings[toTask] -= 1

			if incommings[toTask] == 0 {
				sources.PushBack(toTask)
			}
		}
	}

	return ordering.Len() == tasks
}
