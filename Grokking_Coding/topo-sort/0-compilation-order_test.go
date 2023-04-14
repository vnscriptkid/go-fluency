package topoSort

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	testcases := []struct {
		name           string
		input          [][]rune
		outputExpected []rune
	}{
		{
			name:           "TC1",
			input:          [][]rune{{'B', 'A'}, {'C', 'A'}, {'D', 'C'}, {'E', 'D'}, {'E', 'B'}},
			outputExpected: []rune{'A', 'B', 'C', 'D', 'E'},
		},
		{
			name:           "TC2",
			input:          [][]rune{{'B', 'A'}, {'C', 'A'}, {'D', 'B'}, {'E', 'B'}, {'E', 'D'}, {'E', 'C'}, {'F', 'D'}, {'F', 'E'}, {'F', 'C'}},
			outputExpected: []rune{'A', 'B', 'C', 'D', 'E', 'F'},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			outputActual := findOrder(tc.input)

			if !reflect.DeepEqual(tc.outputExpected, outputActual) {
				t.Errorf("Expected %v but got %v", tc.outputExpected, outputActual)
			}
		})
	}
}
