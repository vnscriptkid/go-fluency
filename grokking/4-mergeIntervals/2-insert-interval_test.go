package grokking

import (
	"reflect"
	"testing"
)

func TestInsertInterval1(t *testing.T) {
	r := insertInterval([]MyInterval{
		{1, 3},
		{5, 7},
		{8, 12},
	}, MyInterval{4, 6})

	if !reflect.DeepEqual(r, []MyInterval{
		{1, 3},
		{4, 7},
		{8, 12},
	}) {
		panic("No match")
	}
}

func TestInsertInterval2(t *testing.T) {
	r := insertInterval([]MyInterval{
		{1, 3},
		{5, 7},
		{8, 12},
	}, MyInterval{4, 10})

	if !reflect.DeepEqual(r, []MyInterval{
		{1, 3},
		{4, 12},
	}) {
		panic("No match")
	}
}

func TestInsertInterval3(t *testing.T) {
	r := insertInterval([]MyInterval{
		{2, 3},
		{5, 7},
	}, MyInterval{1, 4})

	if !reflect.DeepEqual(r, []MyInterval{
		{1, 4},
		{5, 7},
	}) {
		panic("No match")
	}
}
