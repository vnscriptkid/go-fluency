package UnionFindPkg

import (
	"reflect"
	"testing"
)

func Test_UnionFindImpl(t *testing.T) {
	us := MakeUnionFind(5)

	us.Unify(0, 1)

	expectedIDs := []int{0, 0, 2, 3, 4}

	if !reflect.DeepEqual(us.IDs, expectedIDs) {
		t.Errorf("Expected %v but got %v", expectedIDs, us.IDs)
	}

	us.Unify(2, 3)

	expectedIDs = []int{0, 0, 2, 2, 4}

	if !reflect.DeepEqual(us.IDs, expectedIDs) {
		t.Errorf("Expected %v but got %v", expectedIDs, us.IDs)
	}

	us.Unify(3, 4)

	expectedIDs = []int{0, 0, 2, 2, 2}

	if !reflect.DeepEqual(us.IDs, expectedIDs) {
		t.Errorf("Expected %v but got %v", expectedIDs, us.IDs)
	}

	us.Unify(3, 1)

	expectedIDs = []int{2, 0, 2, 2, 2}
	expectedSizes := []int{0, 0, 5, 0, 0}

	if !reflect.DeepEqual(us.IDs, expectedIDs) {
		t.Errorf("Expected %v but got %v", expectedIDs, us.IDs)
	}

	if !reflect.DeepEqual(us.Sizes, expectedSizes) {
		t.Errorf("Expected %v but got %v", expectedSizes, us.Sizes)
	}

	root := us.Find(1)

	if root != 2 {
		t.Errorf("Expected %v but got %v", 2, root)
	}

	expectedIDs = []int{2, 2, 2, 2, 2}

	if !reflect.DeepEqual(us.IDs, expectedIDs) {
		t.Errorf("Expected %v but got %v", expectedIDs, us.IDs)
	}
}
