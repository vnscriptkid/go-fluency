package CustomDS

type SnapshotArray struct {
	// Write your code here
	Snap    [][]int
	Current []int
	Length  int
}

// Constructor intializes the SnapshotArray type object
func InitSnapshotArray(length int) SnapshotArray {
	// Your code will replace this placeholder return statement
	s := SnapshotArray{
		Current: make([]int, length),
		Snap:    make([][]int, 0),
		Length:  length,
	}
	return s
}

// SetValue sets the value at a given index idx to val.
func (s *SnapshotArray) SetValue(idx int, val int) {
	if idx < 0 || idx >= s.Length {
		panic("invalid idx")
	}

	s.Current[idx] = val
}

// Snapshot function takes no parameters and returns the snapid.
// snapid is the number of times that the snapshot() function was called minus 1.
func (s *SnapshotArray) Snapshot() int {
	// Your code will replace this placeholder return statement
	cloned := make([]int, len(s.Current))

	copy(cloned, s.Current)

	s.Snap = append(s.Snap, cloned)

	return len(s.Snap) - 1
}

// GetValue returns the value at the index idx with the given snapid.
func (s *SnapshotArray) GetValue(idx int, snapshotId int) int {
	// Your code will replace this placeholder return statement
	if snapshotId < 0 || snapshotId >= len(s.Snap) {
		panic("invalid snapId")
	}

	return s.Snap[snapshotId][idx]
}
