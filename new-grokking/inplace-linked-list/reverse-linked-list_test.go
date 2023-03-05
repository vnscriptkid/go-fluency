package inplacelinkedlist1

import "testing"

func Test_reverse(t *testing.T) {
	testcases := []struct {
		name           string
		inputRawLL     []int
		outputExpected string
	}{
		{
			name:           "TC1",
			inputRawLL:     []int{1, -2, 3, 4, -5, 4, 3, -2, 1},
			outputExpected: "[1, -2, 3, 4, -5, 4, 3, -2, 1]",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ll := EduLinkedList{}

			ll.CreateLinkedList(tc.inputRawLL)

			outputActual := reverse(ll.head)

			reversedLL := EduLinkedList{
				head: outputActual,
			}

			printed := reversedLL.GetLinkedListPrinted()

			if printed != tc.outputExpected {
				t.Errorf("Expected %v but got %v", tc.outputExpected, printed)
			}
		})
	}
}
