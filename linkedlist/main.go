package main

import "fmt"

func main() {
	// Main: InsertLast, DeleteLast, and Print.
	list := SinglyLinkedList{}
	list.InsertLast("apple")
	list.InsertLast("banana")
	list.InsertLast("cherry")
	list.Print() // Output: apple -> banana -> cherry

	list.DeleteLast()
	list.Print() // Output: apple -> banana

	list.DeleteLast()
	list.Print() // Output: apple

	list.InsertLast("pear")
	list.Print() // Output: apple -> pear

	list.DeleteLast()
	list.Print() // Output: apple

	list.DeleteLast()
	list.Print() // Output: nothing

	list.DeleteLast() // Attempting to delete from an empty list.
	list.Print()      // Output: nothing

	// Optional: Reverse the list.
	fmt.Println("---------------- Reverse:")

	list = SinglyLinkedList{}
	list.InsertLast("apple")
	list.InsertLast("banana")
	list.InsertLast("cherry")
	list.Print() // Output: apple -> banana -> cherry
	list.Reverse()
	list.Print() // Output: cherry -> banana -> apple

	// Optional: InsertFirst.
	fmt.Println("---------------- InsertFirst:")

	list = SinglyLinkedList{}
	list.InsertFirst("banana")
	list.InsertFirst("cherry")
	list.Print() // Output: cherry -> banana

	list.InsertFirst("apple")
	list.Print() // Output: apple -> cherry -> banana

	list.InsertFirst("avocado")
	list.Print() // Output: avocado -> apple -> cherry -> banana

	// Optional: DeleteFirst.
	fmt.Println("---------------- DeleteFirst:")

	list.Print() // Output: avocado -> apple -> cherry -> banana

	list.DeleteFirst()
	list.Print() // Output: apple -> cherry -> banana

	list.DeleteFirst()
	list.Print() // Output: cherry -> banana

	list.DeleteFirst()
	list.Print() // Output: banana

	list.DeleteFirst()
	list.Print() // Output: nothing

	list.DeleteFirst() // Attempting to delete from an empty list.
	list.Print()       // Output: nothing

	// Optional: InsertAt.
	fmt.Println("---------------- InsertAt:")

	list = SinglyLinkedList{}
	list.InsertLast("apple")
	list.InsertLast("banana")
	list.InsertLast("cherry")
	list.Print() // Output: apple -> banana -> cherry

	// Insert "pear" after "banana".
	previous := list.head.next // We know that "banana" is the second node, but in a real-world scenario, we would have to traverse the list to find it or have an 'IndexOf("banana")' kind of method.
	list.InsertAt(previous, "pear")

	list.Print() // Output: apple -> banana -> pear -> cherry

	// Optional: DeleteAt.
	fmt.Println("---------------- DeleteAt:")

	list.Print() // Output: apple -> banana -> pear -> cherry

	// Delete "pear".
	nodeToDelete := list.head.next.next // We know that "pear" is the third node.
	list.DeleteAt(nodeToDelete)

	list.Print() // Output: apple -> banana -> cherry
}
