package QuickUnion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DisjoinSet(t *testing.T) {
	t.Run("basic use case", func(t *testing.T) {
		dSet := NewDSet(10)

		// 1-2-5-6-7 3-8-9 4
		dSet.union(1, 2)
		dSet.union(2, 5)
		dSet.union(5, 6)
		dSet.union(6, 7)
		dSet.union(3, 8)
		dSet.union(8, 9)

		assert.True(t, dSet.connected(1, 5))  // true
		assert.True(t, dSet.connected(5, 7))  // true
		assert.False(t, dSet.connected(4, 9)) // false
		// 1-2-5-6-7 3-8-9-4
		dSet.union(9, 4)
		assert.True(t, dSet.connected(4, 9)) // true

	})
}
