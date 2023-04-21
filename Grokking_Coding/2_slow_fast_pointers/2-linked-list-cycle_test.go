package SlowFastPoitners

import "testing"

func Test_detectCycle(t *testing.T) {

	testcases := []struct {
		InputHead      func() *LinkedListNode
		OutputExpected bool
	}{
		{
			InputHead: func() *LinkedListNode {
				ll := LinkedList{}
				ll.CreateLinkedList([]int{1, 2, 3, 4, 5})
				return ll.head
			},
			OutputExpected: false,
		},
		{
			InputHead: func() *LinkedListNode {
				ll := LinkedList{}
				ll.CreateLinkedList([]int{1, 2, 3, 4, 5})

				node3 := ll.head.next.next
				node5 := ll.head.next.next.next.next

				node5.next = node3

				return ll.head
			},
			OutputExpected: true,
		},
	}

	for _, tc := range testcases {
		head := tc.InputHead()
		output := detectCycle(head)
		if output != tc.OutputExpected {
			t.Errorf("detectCycle(%v) = %v, but expected %v", head, output, tc.OutputExpected)
		}
	}
}
