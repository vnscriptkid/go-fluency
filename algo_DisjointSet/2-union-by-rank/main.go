package UnionByRank

type IDisjointSet interface {
	union(x int, y int)
	find(x int) int
	connected(x int, y int) bool
	rankOf(x int) int
}

type dSet struct {
	root []int
	rank []int
}

func NewDSet(size int) IDisjointSet {
	root := make([]int, size)
	rank := make([]int, size)

	for i := range root {
		// Initially, root of each node is itself, rank is 1
		root[i] = i
		rank[i] = 1
	}

	return &dSet{
		root: root,
		rank: rank,
	}
}

// Time:
func (s *dSet) union(x int, y int) {
	rX := s.root[x]
	rY := s.root[y]

	// Already connected
	if rX == rY {
		return
	}

	if s.rank[rX] == s.rank[rY] {
		s.root[rY] = rX
		s.rank[rX] += 1
	} else if s.rank[rX] > s.rank[rY] {
		s.root[rY] = rX
	} else {
		s.root[rX] = rY
	}
}

// Time: O(n)
func (s *dSet) find(x int) int {
	for x != s.root[x] {
		x = s.root[x]
	}

	return x
}

// Time: O(N)
func (s *dSet) connected(x int, y int) bool {
	return s.find(x) == s.find(y)
}

func (s *dSet) rankOf(x int) int {
	return s.rank[x]
}
