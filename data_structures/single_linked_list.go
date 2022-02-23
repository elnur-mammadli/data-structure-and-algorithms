package main

import (
	"fmt"
)

func main() {
	l := SingleLinkedList{}
	l.insertAtEnd(1)
	l.insertAtEnd(3)
	l.insertAtEnd(4)
	l.insertAtEnd(5)
	l.insertAtEnd(6)
	l.Display()
	n := l.middleNode()
	fmt.Println(n)
}

type Node struct {
	value interface{}
	next  *Node
}

type SingleLinkedList struct {
	head *Node
}

func (L *SingleLinkedList) insertAtHead(value interface{}) {
	n := &Node{
		value: value,
		next: L.head,
	}
	L.head = n
}

func (L *SingleLinkedList) insertAtEnd(value interface{}) {
	n := &Node{
		value: value,
	}

	if L.head == nil {
		L.head = n
		return
	}

	tail := L.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = n
}

func (L *SingleLinkedList) insertAtMiddle(value interface{}, index int) {
	n := &Node{
		value: value,
	}
	if index < 0 {
		return
	}

	if index == 0 {
		n.next = L.head
		return
	}
	idx := 0
	head := L.head

	for head.next != nil {
		if idx + 1 == index {
			head.next = n
			n.next = head.next
			return
		}
		idx++
		head = head.next
	}
}

func (L *SingleLinkedList) Display() {
	head := L.head
	for head != nil {
		fmt.Printf("%v->", head.value)
		head = head.next
	}
	fmt.Println()
}

func (L *SingleLinkedList) RemoveDuplicatesFromLinkedList() {
	head := L.head
	for head.next != nil {
		if head.value == head.next.value {
			head.next = head.next.next
		} else {
			head = head.next
		}
	}
}

func (L *SingleLinkedList) reverse() {
	current := L.head
	var prev *Node
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	L.head = prev
}

func (L *SingleLinkedList) removeElements(val int) {
	for L.head != nil && L.head.value == val { L.head = L.head.next }
	current := L.head
	for current.next != nil {
		if current.next.value == val {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
}

func (L *SingleLinkedList) middleNode() *Node {
	linkMap := make(map[int]*Node)
	current := L.head
	var k int
	for current != nil {
		linkMap[k] = current
		k++
		current = current.next
	}
	lenMap := len(linkMap)
	return linkMap[lenMap / 2]
}