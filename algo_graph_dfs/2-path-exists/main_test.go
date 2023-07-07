package PathExists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validPath(t *testing.T) {
	tcs := []struct {
		name   string
		n      int
		edges  [][]int
		src    int
		dst    int
		expect bool
	}{
		{
			name:   "TC1",
			n:      3,
			edges:  [][]int{{0, 1}, {1, 2}, {2, 0}},
			src:    0,
			dst:    2,
			expect: true,
		},
		{
			name:   "TC2",
			n:      6,
			edges:  [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
			src:    0,
			dst:    5,
			expect: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := validPath(tc.n, tc.edges, tc.src, tc.dst)

			assert.Equal(t, tc.expect, actual)
		})
	}
}
