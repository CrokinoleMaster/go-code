package main

import (
	"errors"
	"fmt"
	"strconv"
)

type integer int

func main() {
	s := new(Stack)
	s.Push(10)
	s.Push(6)
	s.Push(3)
	fmt.Println(s.Min())
}

// integer stack implementation using linkedlist
type Stack struct {
	top *node
	min int
}

type node struct {
	next *node
	data int
}

// returns the minimum integer in the stack
func (s *Stack) Min() int {
	return s.min
}

func (s *Stack) String() string {
	var res string
	for n := s.top; n != nil; n = n.next {
		if n.next != nil {
			res += strconv.Itoa(n.data) + ", "
		} else {
			res += strconv.Itoa(n.data)
		}
	}
	return res
}

func (s *Stack) Push(i int) *Stack {
	if s.top == nil {
		s.min = i
	} else {
		if i < s.min {
			s.min = i
		}
	}
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
