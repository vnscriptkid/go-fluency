package QuickUnion

type IDisjointSet interface {
	union(x int, y int)
	find(x int) int
	connected(x int, y int) bool
}

type dSet struct {
	root []int
}

func NewDSet(size int) IDisjointSet {
	r := make([]int, size)

	for i := range r {
		// Initially, root of each node is itself
		r[i] = i
	}

	return &dSet{
		root: r,
	}
}

// Time: O(N)
func (s *dSet) union(x int, y int) {
	rX := s.find(x)
	rY := s.find(y)

	if rX != rY {
		// Plug root of y to x
		s.root[rY] = x
	}
}

// Time: O(N) in case of skewed tree
func (s *dSet) find(x int) int {
	if s.root[x] == x {
		return x
	}

	return s.find(s.root[x])
}

// Time: O(N)
func (s *dSet) connected(x int, y int) bool {
	return s.find(x) == s.find(y)
}
