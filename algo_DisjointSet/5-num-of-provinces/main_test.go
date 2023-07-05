package NumOfProvinces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCircleNum(t *testing.T) {
	testcases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name:   "TC1",
			input:  [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			expect: 2,
		},
		{
			name:   "TC2",
			input:  [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			expect: 3,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findCircleNum(tc.input)

			assert.Equal(t, tc.expect, actual)
		})
	}
}
