package cache

import "testing"

func TestDLL(t *testing.T) {
	// Create a new DLL with a single node
	dll := NewDLL()
	node := &DllNode{
		Key: "key1",
	}

	dll.AddToFront(node)

	node = &DllNode{
		Key: "key2",
	}

	dll.AddToFront(node)

	dll.Print() // Output: key2->key1->
}
