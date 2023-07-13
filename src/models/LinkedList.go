package models

import (
	"fmt"
)

// Node represents a node in the linked list
type Node struct {
	data interface{}
	next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// Append adds a new node to the end of the linked list
func (list *LinkedList) Append(data interface{}) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Detach removes a node from the linked list
func (list *LinkedList) Detach(value interface{}) {
	if list.head == nil {
		return
	}

	if list.head.data == value {
		list.head = list.head.next
		return
	}

	current := list.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// Search searches for a value in the linked list
func (list *LinkedList) Search(value interface{}) bool {
	current := list.head

	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}

	return false
}

// Size returns the size of the linked list
func (list *LinkedList) Size() int {
	count := 0
	current := list.head

	for current != nil {
		count++
		current = current.next
	}

	return count
}

// Print prints the elements of the linked list
func (list *LinkedList) Print() {
	current := list.head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// DetachHead detaches the first node from the linked list
func (list *LinkedList) DetachHead() {
	if list.head == nil {
		return
	}

	list.head = list.head.next
}

func (list *LinkedList) GetHeadDataAndMoveNext() string { // change implementation of GetHeadDataAndMoveNext
	current := list.head
	list.head = list.head.next
	return current.data.(string)
}

// ConvertToSlice converts the linked list to a slice
func (list *LinkedList) ConvertToSlice() []string {
	slice := make([]string, 0)

	curr := list.head
	for curr != nil {
		if str, ok := curr.data.(string); ok {
			slice = append(slice, str)
		}
		curr = curr.next
	}

	return slice
}
