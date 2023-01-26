package grokking

import (
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/grokking/shared"
)

func Test_reverseAlternatingSublist(t *testing.T) {
	head := shared.ParseLinkedList("1->2->3->4->5->6->7->8->nil")

	newHead := reverseAlternatingSublist(head, 2)

	expected := "2->1->3->4->6->5->7->8"

	if newHead.GetPrinted() != expected {
		t.Errorf("Expected %v but got %v", expected, newHead.GetPrinted())
	}
}
