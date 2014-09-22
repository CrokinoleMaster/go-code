package linkedList

import (
	"errors"
	"strconv"
)

// remove nodes with duplicate values
func (l *LinkedList) RemoveDups() *LinkedList {
	for n := l.root.next; n != l.root; n = n.next {
		for m := l.root.next; m != l.root; m = m.next {
			if m.data == n.data && m != n {
				l.RemoveNode(n)
			}
		}
	}
	return l
}

// find kth to last without using lenth
func (l *LinkedList) FromBack(i int) (*Node, error) {
	if l.len <= i {
		return nil, errors.New("index out of bounds")
	}
	count := 0
	var m *Node
	for n := l.root.next; n != l.root; n = n.next {
		if count == i {
			m = l.root.next
		}
		if count > i {
			m = m.next
		}
		count++
	}
	return m, nil
}

// partition a linked list around value i so that all nodes less than
// i comes before all nodes greater or equal to i
func (l *LinkedList) Partition(i int) *LinkedList {
	for n := l.root.next; n != l.root; {
		if n.data < i {
			if n.next == l.root {
				l.root = n
			}
			m := n.next
			n.next = m.next
			n.data, m.data = m.data, n.data
			m.next = l.root.next
			l.root.next = m
		} else {
			n = n.next
		}
	}
	return l
}

// delete node without head access
// **does not update linkedlist length**
func DeleteNode(n *Node) error {
	m := n.next
	if m == nil {
		return errors.New("node is not in a list")
	}
	n.next = m.next
	n.data = m.data
	return nil
}

// add two numbers represented by a linked list
// ex) (7-> 1 -> 6) + (5 -> 9 -> 2) = (2 -> 1 -> 9)
//       617		+	295			=	912
func AddDecimalList(l, _l *LinkedList) *LinkedList {
	res := NewLinkedList()
	length := l.len
	_length := _l.len
	var long *LinkedList
	var short *LinkedList
	if length > _length {
		long = l
		short = _l
	} else {
		long = _l
		short = l
	}

	for n, sn := long.root.next, short.root.next; n != long.root; n = n.next {
		if sn != short.root {
			sum := n.data + sn.data
			if sum > 9 {
				sum = sum - 10
				if n.next != long.root {
					n.next.data += 1
				} else {
					long.Push(1)
				}
			}
			res.Push(sum)
		} else {
			res.Push(n.data)
		}
		sn = sn.next

	}
	return res
}

// checks if a given linked list is a palindrome
func (l *LinkedList) isPalindrome() bool {
	_l := NewLinkedList()
	prev := _l.root
	for n := l.root.next; n != l.root; n = n.next {
		_n := &Node{data: n.data}
		_l.root.next = _n
		_n.next = prev
		prev = _n
	}
	for n, _n := l.root.next, _l.root.next; n != l.root; n = n.next {
		if n.data != _n.data {
			return false
		}
		_n = _n.next
	}
	return true
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
			l.len--
		}
	}
	return l
}

func (l *LinkedList) RemoveNode(e *Node) *LinkedList {
	if l.len == 0 {
		return l
	}
	for n := l.root.next; n != l.root; n = n.next {
		if n.next == e {
			n.next = n.next.next
			l.len--
		}
	}
	return l
}
