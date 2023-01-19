package grokking

import (
	"container/list"
)

func tasksSchedulingOrder(tasks int, prerequisites [][2]int) []int {
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

	sources := list.New()

	for task, numOfIncommings := range incommings {
		if numOfIncommings == 0 {
			sources.PushBack(task)
		}
	}

	order := []int{}

	for sources.Len() > 0 {
		task := sources.Remove(sources.Front()).(int)

		order = append(order, task)

		for _, destTask := range graph[task] {
			incommings[destTask] -= 1

			if incommings[destTask] == 0 {
				sources.PushBack(destTask)
			}
		}
	}

	if len(order) != tasks {
		return []int{}
	}

	return order
}
