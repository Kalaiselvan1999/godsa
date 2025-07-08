package main

import "github.com/Kalaiselvan1999/godsa/algorithms"


func main() {
	dll := &algorithms.DoublyLinkedList{}
	dll.ArrayToDLL([]int{10, 20, 30, 40})
	dll.Display()

	dll.InsertBeforeHead(5)
	dll.InsertAfterHead(15)
	dll.InsertAfterTail(50)
	dll.InsertBeforeTail(45)
	dll.InsertBeforeKth(2, 12)
	dll.InsertAfterKth(3, 25)

	dll.Display()

	dll.DeleteHead()
	dll.DeleteTail()
	dll.DeleteKthElement(2)
	dll.DeleteNode(25)

	dll.Display()
}