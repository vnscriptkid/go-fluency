package TraverseAllPaths

import (
	"fmt"
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

		toABC := []string{"A", "B", "C", "D", "E", "F"}

		expect := [][]int{{0, 1}, {0, 2, 4, 1}, {0, 2, 4, 5, 1}, {0, 3, 4, 1}, {0, 3, 4, 5, 1}}
		expectABC := [][]string{}

		for _, e := range expect {
			x := []string{}
			for _, i := range e {
				x = append(x, toABC[i])
			}
			expectABC = append(expectABC, x)
		}

		actual := traverse(g, 0, 1)

		assert.Equal(t, expect, actual)

		fmt.Println(expectABC)
	})
}
