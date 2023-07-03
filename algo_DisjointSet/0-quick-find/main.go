package QuickFind

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
	// (1)-2  (3)-4
	// union(2,4)

	rX := s.find(x)
	rY := s.find(y)

	// Only do union if 2 nodes are of different groups
	if rX != rY {
		for i := range s.root {
			if s.root[i] == rY {
				// Connecting 2 groups
				// For each and every node of group to which y belong, make rX as new root
				s.root[i] = rX
			}
		}
	}
}

// Time: O(1)
func (s *dSet) find(x int) int {
	return s.root[x]
}

// Time: O(1)
func (s *dSet) connected(x int, y int) bool {
	return s.find(x) == s.find(y)
}
