package main

import (
	"errors"
	"fmt"
	"strings"
)

// check if all characters are unique
func Unique(s string) bool {
	for i, c := range s {
		for _i, _c := range s {
			if i != _i && c == _c {
				return false
			}
		}
	}
	return true
}

// check if strings are permutations
func Permutation(s, _s string) bool {
	if len(s) != len(_s) {
		return false
	}
	r := make([]bool, len(s))
	for _, c := range s {
		for _i, _c := range _s {
			if c == _c && !r[_i] {
				r[_i] = true
				break
			}
			if _i == len(_s)-1 {
				return false
			}
		}
	}
	return true
}

// encode spaces into %20
func EncodeSpaces(s string) string {
	l := len(s)
	ss := []rune(s)
	r := make([]rune, l*3)
	f := false
	e := l
	for i := l - 1; i >= 0; i-- {
		r[i] = ss[i]
		if r[i] != ' ' {
			if !f {
				f = true
				e = i
			}
		} else if f {
			for j := e; j > i; j-- {
				r[j+2] = r[j]
			}
			r[i] = '%'
			r[i+1] = '2'
			r[i+2] = '0'
			e += 2
		}
	}
	fmt.Println(r[:e+1])
	return string(r[:e+1])
}

// compress string using counts of repeated characters
func Compress(s string) string {
	l := len(s)
	r := make([]rune, l*2)
	var p rune
	j := 0
	count := 0
	for i, c := range s {
		if c != p {
			r[j] = c
			if i != 0 {
				r[j-1] = rune(count + '0')
			}
			j += 2
			if j > l {
				return s
			}
			p = c
			count = 0
		}
		count++
		if i == l-1 {
			fmt.Println(count)
			r[j-1] = rune(count + '0')
		}
	}
	return string(r[:j])
}

// rotate matrix 90 degrees, NxN matrix
func RotateMatrix(m [][]int) ([][]int, error) {
	l := len(m)
	if l != len(m[0]) {
		return m, errors.New("matrix not N by N")
	}
	for ri := 0; ri < l/2; ri++ {
		for ci := ri; ci < l-ri-1; ci++ {
			temp := m[ri][ci]
			m[ri][ci] = m[l-ci-1][ri]
			m[l-ci-1][ri] = m[l-ri-1][l-ci-1]
			m[l-ri-1][l-ci-1] = m[ci][l-ri-1]
			m[ci][l-ri-1] = temp
		}
	}
	return m, nil
}

// if an element is 0, zero out it's row and column
func ZeroOut(m [][]int) [][]int {
	rows := len(m)
	cols := len(m[0])
	ma := make(map[int]bool, rows+cols)
	for ri, row := range m {
		for ci, n := range row {
			if n == 0 {
				ma[-ci] = true
				ma[ri] = true
			}
		}
	}
	for ri, row := range m {
		for ci, _ := range row {
			if ma[-ci] || ma[ri] {
				m[ri][ci] = 0
			}
		}
	}

	return m
}

// checks if two strings are rotations of each other
func IsRotation(s, _s string) bool {
	l := len(s)
	_l := len(_s)
	if l != _l {
		return false
	}
	s = s + s
	if strings.Contains(s, _s) {
		return true
	}

	return false
}

func main() {
}
