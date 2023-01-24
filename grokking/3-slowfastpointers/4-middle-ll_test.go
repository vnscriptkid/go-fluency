package grokking

import (
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/grokking/shared"
)

func TestMiddleOfLinkedList(t *testing.T) {
	testcases := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "TC1",
			input:          "1->2->3->4->5->nil",
			expectedOutput: 3,
		},
		{
			name:           "TC2",
			input:          "1->2->3->4->5->6->nil",
			expectedOutput: 4,
		},
		{
			name:           "TC3",
			input:          "1->2->3->4->5->6->7->nil",
			expectedOutput: 4,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			head := shared.ParseLinkedList(tc.input)

			actual := middleOfLinkedList(head)

			if actual != tc.expectedOutput {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
