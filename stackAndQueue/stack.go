package main

import "errors"

type integer int

func main() {
}

// integer stack implementation using linkedlist
type Stack struct {
	top *node
}

type node struct {
	next *node
	data int
}

func (s *Stack) Push(i int) *Stack {
	n := &node{next: nil, data: i}
	n.next = s.top
	s.top = n
	return s
}

func (s *Stack) Pop(i int) (int, error) {
	if s.top == nil {
		return 0, errors.New("Not found")
	}
	s.top = s.top.next
	return s.top.data, nil
}

func (s *Stack) Peek() int {
	return s.top.data
}
