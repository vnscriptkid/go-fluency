package grokking

import (
	"strconv"
	"strings"
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/grokking/shared"
)

func ParseLinkedList(str string) *shared.Node {
	splitted := strings.Split(str, "->")

	var head *shared.Node
	var prev *shared.Node

	for _, strVal := range splitted {
		if val, err := strconv.Atoi(strVal); err == nil {
			node := shared.Node{Value: val}

			if prev == nil {
				head = &node
			} else {
				prev.Next = &node
			}

			prev = &node
		} else {
			// end of list
		}
	}

	return head
}

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
			head := ParseLinkedList(tc.input)

			actual := middleOfLinkedList(head)

			if actual != tc.expectedOutput {
				t.Errorf("Expected %v but got %v", tc.expectedOutput, actual)
			}
		})
	}
}
