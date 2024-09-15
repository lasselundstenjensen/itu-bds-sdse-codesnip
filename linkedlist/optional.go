package main

// Reverse reverses the linked list in place.
func (sll *SinglyLinkedList) Reverse() {
	// Declare pointers.
	var previous *Node
	current := sll.head
	var next *Node

	for current != nil {
		next = current.next     // Save the next node.
		current.next = previous // Reverse the current node's pointer.
		previous = current      // Move previous to the current node.
		current = next          // Move to the next node.
	}
	sll.head = previous // Update the head to the new first node.
}

// InsertFirst adds a new node with the given data at the beginning of the list.
func (sll *SinglyLinkedList) InsertFirst(data string) {
	new := &Node{data: data}
	new.next = sll.head
	sll.head = new
}

// DeleteFirst removes the first node from the list.
func (sll *SinglyLinkedList) DeleteFirst() {
	if sll.head == nil {
		// The list is already empty.
		return
	}
	sll.head = sll.head.next
}

// InsertAt inserts a new node with the given data after the specified node.
// Note: This is one way to insert a node in a linked list, but there may be other ways to implement it.
func (sll *SinglyLinkedList) InsertAt(previous *Node, data string) {
	if previous == nil {
		panic("the given previous node cannot be nil")
	}
	new := &Node{data: data}
	new.next = previous.next
	previous.next = new
}

// DeleteAt deletes the specified node from the list.
func (sll *SinglyLinkedList) DeleteAt(delete *Node) {
	if sll.head == nil || delete == nil {
		// The list is empty or the node to delete is nil.
		return
	}

	if sll.head == delete {
		// The node to delete is the head.
		sll.head = sll.head.next
		return
	}

	// Find the previous node.
	current := sll.head
	var previous *Node

	// Traverse the list to find the node to delete, while keeping track of the previous node.
	for current != nil && current != delete {
		previous = current
		current = current.next
	}

	if current == nil {
		// The node to delete is not in the list.
		panic("the node to delete is not in the list")
	}

	// Found the node to delete (current). Remove node to delete from the list by updating the pointer.
	previous.next = current.next
}
