package algorithms

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head *Node
}

// Convert array to DLL
func (dll *DoublyLinkedList) ArrayToDLL(arr []int) {
	if len(arr) == 0 {
		return
	}
	dll.Head = &Node{Data: arr[0]}
	current := dll.Head

	for i := 1; i < len(arr); i++ {
		newNode := &Node{Data: arr[i], Prev: current}
		current.Next = newNode
		current = newNode
	}
}

// Display forward
func (dll *DoublyLinkedList) Display() {
	current := dll.Head
	fmt.Print("DLL Display: ")
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

// Delete head node
func (dll *DoublyLinkedList) DeleteHead() {
	if dll.Head == nil {
		return
	}
	dll.Head = dll.Head.Next
	if dll.Head != nil {
		dll.Head.Prev = nil
	}
}

// Delete tail node
func (dll *DoublyLinkedList) DeleteTail() {
	if dll.Head == nil {
		return
	}
	current := dll.Head
	for current.Next != nil {
		current = current.Next
	}
	if current.Prev != nil {
		current.Prev.Next = nil
	} else {
		dll.Head = nil // only one node
	}
}

// Delete k-th (0-indexed) node
func (dll *DoublyLinkedList) DeleteKthElement(index int) {
	if index < 0 || dll.Head == nil {
		return
	}
	current := dll.Head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		return
	}
	if current.Prev != nil {
		current.Prev.Next = current.Next
	} else {
		dll.Head = current.Next
	}
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}
}

// Delete a node by value (first occurrence)
func (dll *DoublyLinkedList) DeleteNode(value int) {
	current := dll.Head
	for current != nil && current.Data != value {
		current = current.Next
	}
	if current == nil {
		return
	}
	if current.Prev != nil {
		current.Prev.Next = current.Next
	} else {
		dll.Head = current.Next
	}
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}
}

// Insert before head
func (dll *DoublyLinkedList) InsertBeforeHead(data int) {
	newNode := &Node{Data: data, Next: dll.Head}
	if dll.Head != nil {
		dll.Head.Prev = newNode
	}
	dll.Head = newNode
}

// Insert before tail
func (dll *DoublyLinkedList) InsertBeforeTail(data int) {
	if dll.Head == nil {
		dll.InsertBeforeHead(data)
		return
	}
	current := dll.Head
	for current.Next != nil {
		current = current.Next
	}
	newNode := &Node{Data: data, Prev: current.Prev, Next: current}
	if current.Prev != nil {
		current.Prev.Next = newNode
	} else {
		dll.Head = newNode
	}
	current.Prev = newNode
}

// Insert before k-th index
func (dll *DoublyLinkedList) InsertBeforeKth(index int, data int) {
	if index <= 0 || dll.Head == nil {
		dll.InsertBeforeHead(data)
		return
	}
	current := dll.Head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		dll.InsertBeforeTail(data)
		return
	}
	newNode := &Node{Data: data, Prev: current.Prev, Next: current}
	if current.Prev != nil {
		current.Prev.Next = newNode
	}
	current.Prev = newNode
	if index == 0 {
		dll.Head = newNode
	}
}

// Insert before specific node
func (dll *DoublyLinkedList) InsertBeforeNode(data int, node *Node) {
	if node == nil {
		return
	}
	newNode := &Node{Data: data, Prev: node.Prev, Next: node}
	if node.Prev != nil {
		node.Prev.Next = newNode
	} else {
		dll.Head = newNode
	}
	node.Prev = newNode
}

// Insert after head
func (dll *DoublyLinkedList) InsertAfterHead(data int) {
	if dll.Head == nil {
		dll.InsertBeforeHead(data)
		return
	}
	newNode := &Node{Data: data, Prev: dll.Head, Next: dll.Head.Next}
	if dll.Head.Next != nil {
		dll.Head.Next.Prev = newNode
	}
	dll.Head.Next = newNode
}

// Insert after tail
func (dll *DoublyLinkedList) InsertAfterTail(data int) {
	if dll.Head == nil {
		dll.InsertBeforeHead(data)
		return
	}
	current := dll.Head
	for current.Next != nil {
		current = current.Next
	}
	newNode := &Node{Data: data, Prev: current}
	current.Next = newNode
}

// Insert after k-th node
func (dll *DoublyLinkedList) InsertAfterKth(index int, data int) {
	if index < 0 {
		return
	}
	current := dll.Head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		dll.InsertAfterTail(data)
		return
	}
	newNode := &Node{Data: data, Prev: current, Next: current.Next}
	if current.Next != nil {
		current.Next.Prev = newNode
	}
	current.Next = newNode
}

// Insert after a given node
func (dll *DoublyLinkedList) InsertAfterNode(data int, node *Node) {
	if node == nil {
		return
	}
	newNode := &Node{Data: data, Prev: node, Next: node.Next}
	if node.Next != nil {
		node.Next.Prev = newNode
	}
	node.Next = newNode
}

func LoadDll() {
	dll := DoublyLinkedList{}
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