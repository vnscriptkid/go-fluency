package subsets

// findAllSubsets will find all the subsets of the given array
func findAllSubsets(v []int) [][]int {
	// Write your code here
	// Your code will replace this placeholder return statement

	// 1,2,3
	// [ [] ]
	//    ^
	// 1: [ [], [1] ]
	//       ^
	// 2: [ [], [1], [2], [1,2] ]
	//                      ^
	// 3: [ [], [1], [2], [1,2], [3], [1,3], [2,3], [1,2,3] ]

	result := make([][]int, 0)

	result = append(result, []int{})

	for _, num := range v {
		levelSize := len(result)

		for i := 0; i < levelSize; i++ {
			curSet := result[i]

			cloneSet := make([]int, len(curSet))

			copy(cloneSet, curSet)

			cloneSet = append(cloneSet, num)

			result = append(result, cloneSet)
		}
	}

	return result
}

type Set struct {
	// We will use map to implement the set.
	// To make map work as a set the type of value in a map is bool.
	hashMap map[int]bool
}

// NewSet will initialize and return the new object of Set.
func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[int]bool)
	return s
}

func NewSetFromArr(arr []int) *Set {
	s := NewSet()

	for _, num := range arr {
		if !s.Exists(num) {
			s.Add(num)
		}
	}

	return s
}

// Add will add the value in the Set.
func (s *Set) Add(value int) {
	s.hashMap[value] = true
}

// Delete will delete the value from the set.
func (s *Set) Delete(value int) {
	delete(s.hashMap, value)
}

// Exists will check if the value exists in the set or not.
func (s *Set) Exists(value int) bool {
	_, ok := s.hashMap[value]
	return ok
}
