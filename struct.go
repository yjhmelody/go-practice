package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Postion   string
	Salary    int
	ManagetID int
}

func main() {
	var yjh Employee
	yjh.Name = "yjh"
	yjh.Salary = 50000

	yjhptr := &yjh
	fmt.Println(yjhptr.Name)
	fmt.Println((*yjhptr).Name)

	var values = []int{5, 6, 3, 1, 2, 4}
	Sort(values)
	for _, v := range values {
		fmt.Println(v)
	}
	fmt.Println(values[:0])

	var p = Point{1, 2}
	fmt.Println(p.x, p.y)

	var p2 = Point{x: 2, y: 4}
	fmt.Println(p2.x, p2.y)
}

type Point struct {
	x int
	y int
}

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
// 前序遍历二叉树
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
