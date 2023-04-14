package CustomDS

import (
	"testing"
)

func Test_SnapshotArr(t *testing.T) {
	snapshotArr := InitSnapshotArray(3)
	snapshotArr.SetValue(0, 6)
	x := snapshotArr.Snapshot()

	if x != 0 {
		t.Errorf("Expect 0 but got %v", x)
	}

	x = snapshotArr.GetValue(0, 0)

	if x != 6 {
		t.Errorf("Expect 6 but got %v", x)
	}

	snapshotArr.SetValue(1, 8)

	x = snapshotArr.Snapshot()

	if x != 1 {
		t.Errorf("Expect 1 but got %v", x)
	}

	x = snapshotArr.GetValue(1, 1)

	if x != 8 {
		t.Errorf("Expect 8 but got %v", x)
	}
}
