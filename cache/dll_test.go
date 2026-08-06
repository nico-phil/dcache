package cache

import (
	"fmt"
	"testing"
)

func TestDLL(t *testing.T) {
	// Create a new DLL with a single node
	dll := NewDLL()

	for i := 0; i < 5; i++ {
		newNode := &DllNode{
			Key: fmt.Sprintf("key%d", i+1),
		}
		dll.AddToFront(newNode)
	}

	dll.Print() // Output: key2->key1->

	dll.Remove(dll.Tail)
	dll.Print()

	dll.MoveToFront(dll.Tail)
	dll.Print()
}
