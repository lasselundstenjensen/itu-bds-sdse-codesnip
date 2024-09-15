package main

import "fmt"

// Node represents a node in the linked list.
type Node struct {
	data string
	next *Node
}

// SinglyLinkedList represents the linked list.
type SinglyLinkedList struct {
	head *Node
}

// InsertLast adds a new node with the given string to the end of the list.
func (sll *SinglyLinkedList) InsertLast(data string) {
	new := &Node{data: data}

	if sll.head == nil {
		sll.head = new
		return
	}

	previous := sll.head
	for previous.next != nil {
		previous = previous.next
	}
	previous.next = new
}

// DeleteLast removes the last node from the list and updates the head if necessary.
func (sll *SinglyLinkedList) DeleteLast() {
	// Guard clause:
	if sll.head == nil {
		// The list is already empty.
		return
	}

	// Guard clause:
	if sll.head.next == nil {
		// Only one node in the list.
		sll.head = nil
		return
	}

	// Declare pointers.
	current := sll.head
	var previous *Node

	// Traverse to the last node.
	for current.next != nil {
		previous = current
		current = current.next
	}

	// Remove the last node.
	previous.next = nil
}

// Print displays all the elements of the list.
// Example: headValue -> node1Value -> node2Value -> node3Value
func (sll *SinglyLinkedList) Print() {
	current := sll.head
	for current != nil {
		fmt.Printf("%s", current.data)
		current = current.next
		if current != nil {
			fmt.Printf(" -> ")
		}
	}
	fmt.Println()
}
