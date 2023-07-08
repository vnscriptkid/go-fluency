package ShortestPath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_traverse(t *testing.T) {
	t.Run("TC 1", func(t *testing.T) {
		expect := 2

		edges := [][2]string{
			{"A", "C"},
			{"A", "D"},
			{"A", "B"},
			{"C", "E"},
			{"D", "E"},
			{"B", "E"},
			{"B", "F"},
			{"E", "F"},
		}

		actual := findShortestPath(edges, "A", "F")

		assert.Equal(t, expect, actual)
	})
}
