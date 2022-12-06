package main

import (
	"container/list" // Doubly-linkedlist
	"fmt"
)

func main() {
	// Going forward from head
	var strList list.List

	strList.PushBack("one")
	strList.PushBack("two")
	strList.PushBack("three")

	for ele := strList.Front(); ele != nil; ele = ele.Next() {
		// fmt.Println("Ele: ", ele.Value)

		// TODO: What the diff?
		// eleValue := string (ele.Value); <<<< Can not cast

		// TODO: What is this called?
		// TODO: Type `any` means what?
		eleValue := ele.Value.(string)

		fmt.Println("Ele: ", eleValue)

	}

	// Going backward from rear
	var intList list.List

	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)

	for ele := intList.Back(); ele != nil; ele = ele.Prev() {
		fmt.Println("Int Ele: ", ele.Value.(int))
	}
}
