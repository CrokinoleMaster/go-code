package main

import (
	"errors"
	"fmt"

	"github.com/huaruiwu/go-code/stack"
)

func main() {
	s := new(stack.Stack)
	s.Push(1)
	s.Push(3)
	s.Push(1)
	s.Push(2)
	s.Push(2)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	n := stack.SortStack(*s)
	fmt.Println(n)
	n.Pop()
	n.Pop()
	fmt.Println(n)
}

// queue structure using two stacks
type Queue struct {
	in  *stack.Stack
	out *stack.Stack
}

func NewQueue() *Queue {
	return &Queue{new(stack.Stack), new(stack.Stack)}
}

func (q *Queue) Queue(i int) {
	q.in.Push(i)
}

func (q *Queue) Deque() (int, error) {
	n, err := q.out.Pop()
	if err == nil {
		return n, nil
	}
	for n, err := q.in.Pop(); err == nil; n, err = q.in.Pop() {
		q.out.Push(n)
	}
	n, err = q.out.Pop()
	if err != nil {
		return 0, errors.New("empty queue")
	}
	return n, nil
}
