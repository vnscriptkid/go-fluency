package grokking

import "testing"

type Args struct {
	arr []int
	sum int
}

func TestSmallestSubArrWithSum(t *testing.T) {
	testcases := []struct {
		name     string
		args     Args
		expected int
	}{
		{
			name:     "TC1",
			args:     Args{arr: []int{2, 1, 5, 2, 3, 2}, sum: 7},
			expected: 2,
		},
		{
			name:     "TC2",
			args:     Args{arr: []int{2, 1, 5, 2, 8}, sum: 7},
			expected: 1,
		},
		{
			name:     "TC3",
			args:     Args{arr: []int{3, 4, 1, 1, 6}, sum: 8},
			expected: 3,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			r := smallestSubArrWithSum(tc.args.arr, tc.args.sum)

			if r != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, r)
			}
		})
	}
}
