package main

import (
	"fmt"
	"strconv"
)

func main() {
	l := NewLinkedList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	l.Push(6)
	l.Push(7)
	l.Push(8)
	l.Remove(8)
	l.Remove(4)

	fmt.Println(l)
}

// integer linked list
type LinkedList struct {
	root *Node
	len  int
}

type Node struct {
	next *Node
	data int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{&Node{}, 0}
}

func (l *LinkedList) String() string {
	var s string
	for n := l.root.next; n != l.root; n = n.next {
		s += strconv.Itoa(n.data)
		if n.next != l.root {
			s += ", "
		}
	}
	return s
}

func (e *Node) Next() *Node {
	n := e.next
	return n
}

func (l *LinkedList) Clear() *LinkedList {
	l.root.next = l.root
	l.len = 0
	return l
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *LinkedList) Push(i int) *Node {
	var n *Node
	if l.len == 0 {
		n = &Node{data: i}
		n.next = l.root
		l.root.next = n
		l.len++
		return n
	}
	for n = l.root.next; n != l.root; n = n.next {
		if n.next == l.root {
			n.next = &Node{data: i}
			n.next.next = l.root
			l.len++
			return n
		}
	}
	return nil
}

func (l *LinkedList) Find(i int) *Node {
	if l.len == 0 {
		return nil
	}
	for n := l.root.next; n != l.root; n = n.next {
		if n.data == i {
			return n
		}
	}
	return nil
}

func (l *LinkedList) Remove(i int) *LinkedList {
	for n := l.root.next; n != l.root; n = n.next {
		if n.next.data == i {
			n.next = n.next.next
		}
	}
	return l
}
