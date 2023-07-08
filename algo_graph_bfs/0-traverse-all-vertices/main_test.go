package TraverseAllVertices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_traverse(t *testing.T) {
	t.Run("TC 1", func(t *testing.T) {
		g := [][]int{
			//   A  B  C  D  E  F
			{1, 1, 1, 1, 0, 0}, // A
			{1, 1, 0, 0, 1, 1}, // B
			{1, 0, 1, 0, 1, 0}, // C
			{1, 0, 0, 1, 1, 0}, // D
			{0, 1, 1, 1, 1, 1}, // E
			{0, 1, 0, 0, 1, 1}, // F
		}

		expect := []int{0, 1, 2, 3, 4, 5}

		actual := traverse(g)

		assert.Equal(t, expect, actual)
	})
}
