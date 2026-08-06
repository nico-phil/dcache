package cache

import "fmt"

type DLL struct {
	Head *DllNode
	Tail *DllNode
}

type DllNode struct {
	Key string
	// Value []byte
	Prev *DllNode
	Next *DllNode
}

func NewDLL() *DLL {
	return &DLL{}
}

func (dll *DLL) AddToFront(node *DllNode) {
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		return
	}

	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
}

func (dll *DLL) Remove(node *DllNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		dll.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		dll.Tail = node.Prev
	}
}

func (dll *DLL) MoveToFront(node *DllNode) {
	dll.Remove(node)
	dll.AddToFront(node)
}

func (dll *DLL) Print() {
	current := dll.Head
	for current != nil {
		fmt.Printf("%s->", current.Key)
		current = current.Next
	}
	fmt.Println("\nTail:", dll.Tail.Key)
}
