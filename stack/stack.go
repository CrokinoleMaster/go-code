package stack

import (
	"errors"
	"fmt"
	"strconv"
)

// sort stack in ascending order using a temp stack
func SortStack(s Stack) *Stack {
	ss := new(Stack)
	for i, n := s.Pop(); n == nil; i, n = s.Pop() {
		count := 0
		for t, err := ss.Peek(); i < t; t, err = ss.Peek() {
			if err == nil {
				n, _ := ss.Pop()
				s.Push(n)
				count++
			}
		}
		ss.Push(i)
		for ; count > 0; count-- {
			i, _ := s.Pop()
			ss.Push(i)
		}
	}
	return ss
}

// a data structure that starts a new stack when the previous stack exceeds
// a given threshold
type SetOfStacks struct {
	limit     int
	currStack *Stack
}

func NewSetOfStacks(limit int) *SetOfStacks {
	s := &SetOfStacks{limit: limit, currStack: new(Stack)}
	return s
}

func (s *SetOfStacks) Push(i int) *SetOfStacks {
	n := &node{next: s.currStack.top, data: i}
	if s.currStack.length == s.limit {
		st := new(Stack)
		st.next = s.currStack
		s.currStack = st
		s.currStack.length = 0
	}
	s.currStack.top = n
	s.currStack.length += 1
	return s
}

func (s *SetOfStacks) Peek() (int, error) {
	if s.currStack == nil {
		return 0, errors.New("empty stack")
	}
	return s.currStack.top.data, nil
}

func (s *SetOfStacks) Pop() (int, error) {
	if s.currStack.top == nil {
		return 0, errors.New("empty stack")
	}
	res, err := s.currStack.Pop()
	if err != nil {
		fmt.Println("Pop error")
	}
	if s.currStack.length == 0 {
		s.currStack = s.currStack.next
	}
	return res, nil
}

// integer stack implementation using linkedlist
type Stack struct {
	top    *node
	next   *Stack
	min    int
	length int
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
	s.length += 1
	return s
}

func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return 0, errors.New("empty stack")
	}
	res := s.top.data
	s.top = s.top.next
	s.length -= 1
	return res, nil
}

func (s *Stack) Peek() (int, error) {
	if s.top == nil {
		return 0, errors.New("empty stack")
	}
	return s.top.data, nil
}
