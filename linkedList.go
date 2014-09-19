package main

// integer linked list
type LinkedList struct {
	root Node
	len  int
}

type Node struct {
	next *Node
	data int
}

func (e *Node) Next() *Node {
	n := e.next
	return n
}

func (l *LinkedList) Clear() *LinkedList {
	l.root.next = nil
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
	for n = l.root.next; n != nil; n = n.next {
		if n.next == nil {
			n.next = &Node{data: i}
		}
	}
	return n
}

func (l *LinkedList) Find(i int) *Node {
	for n := l.root.next; n.next != nil; n = n.next {
		if n.next.data == i {
			return n
		}
	}
	return nil
}

func (l *LinkedList) Remove(i int) *LinkedList {
	n := l.Find(i)
	if n == nil {
		return nil
	}
	n.next = n.next.next
	l.len--
	return l
}

func main() {

}
