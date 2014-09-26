package main

import "fmt"

func main() {
	b := &BST{5, nil, nil}
	Insert(b, 1)
	Insert(b, 2)
	Insert(b, 3)
	Insert(b, 4)
	Insert(b, 6)
	Insert(b, 7)
	Insert(b, 8)
	Insert(b, 9)
	fmt.Println(b.Right)
}

// binary tree
type BST struct {
	data        int
	Left, Right *BST
}

func Insert(b *BST, i int) *BST {
	if b == nil {
		return &BST{i, nil, nil}
	}
	if i < b.data {
		b.Left = Insert(b.Left, i)
	} else {
		b.Right = Insert(b.Right, i)
	}
	return b
}

func Contains(b *BST, i int) bool {
	if b == nil {
		return false
	} else if b.data == i {
		return true
	}
	if i < b.data {
		return Contains(b.Left, i)
	} else {
		return Contains(b.Right, i)
	}
}
