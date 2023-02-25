package grokking

func dfs(graph map[int][]int, incommings map[int]int, candidates []int, cur *[]int, all *[][]int) {
	if len(*cur) == len(graph) {
		cloneCur := make([]int, len(*cur))
		copy(cloneCur, *cur)
		*all = append(*all, cloneCur)
		return
	}

	for i, can := range candidates {
		*cur = append(*cur, can)

		newCandidates := make([]int, 0, len(candidates)-1)
		newCandidates = append(newCandidates, candidates[:i]...)
		newCandidates = append(newCandidates, candidates[i+1:]...)

		for _, dest := range graph[can] {
			incommings[dest] -= 1

			if incommings[dest] == 0 {
				newCandidates = append(newCandidates, dest)
			}
		}

		dfs(graph, incommings, newCandidates, cur, all)

		*cur = (*cur)[0 : len(*cur)-1]

		for _, dest := range graph[can] {
			incommings[dest] += 1
		}
	}
}

func allTasksSchedulingOrder(tasks int, prerequisites [][2]int) [][]int {
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

	candidates := []int{}

	for task, numOfIncommings := range incommings {
		if numOfIncommings == 0 {
			candidates = append(candidates, task)
		}
	}

	// {3, 2}, {3, 0}, {2, 0}, {2, 1}
	// graph
	// 0 : []
	// 1 : []
	// 2 : [0,1]
	// 3 : [0,2]
	// incommings
	// 0 : 0
	// 1 : 0
	// 2 : 0
	// 3 : 0

	// 0,1
	//
	// sources: []
	// 2
	// [3,2]

	all := [][]int{}

	dfs(graph, incommings, candidates, &[]int{}, &all)

	return all
}
