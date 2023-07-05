package NumOfProvinces

/*
Thought process: connected cities make a union (set of connected cities)
We can represent provinces as disjoint sets
(1) Build disjoint set using array
// [ 1, 1, 0 ] parent
//   0  1  2
// parent[i]: the parent of city i

(2)How to count provinces:
- Loop through all cities, for each we find it's root in the union
- The number of provinces is the number of roots
- We can keep distinct roots in a set
- At the end, we count number of root in set
*/

type IDisjointSet interface {
	find(x int) int
	union(x int, y int)
	connected(x int, y int) bool
}

type dSet struct {
	parent []int
	rank   []int // Avoid skewed tree
}

func (s *dSet) find(x int) int {
	if x == s.parent[x] {
		return x
	}

	// Improvement: Path compression
	return s.find(s.parent[x])
}

func (s *dSet) union(x int, y int) {
	rX := s.find(x)
	rY := s.find(y)

	if rX == rY {
		// Already connected
		return
	}

	// Case #1
	if s.rank[rX] == s.rank[rY] {
		// 1 <- 2
		// ^
		// 3 <- 4
		s.parent[rY] = rX
		s.rank[rX] += 1
	} else if s.rank[rX] > s.rank[rY] {
		// Case #2
		// 1 <- 2
		// 3
		s.parent[rY] = rX
	} else {
		// Case #3
		s.parent[rX] = rY
	}
}

func (s *dSet) connected(x int, y int) bool {
	return s.find(x) == s.find(y)
}

func NewDSet(size int) IDisjointSet {
	parent := make([]int, size)
	rank := make([]int, size)

	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}

	return &dSet{
		parent: parent,
		rank:   rank,
	}
}

func findCircleNum(m [][]int) int {
	n := len(m)

	s := NewDSet(n)

	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if m[i][j] == 1 {
				s.union(i, j)
			}
		}
	}

	seen := map[int]bool{}

	for i := 0; i < n; i++ {
		r := s.find(i)

		if ok := seen[r]; !ok {
			seen[r] = true
		}
	}

	return len(seen)
}
