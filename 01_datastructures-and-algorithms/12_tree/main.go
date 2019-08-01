package main

import "fmt"

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func (tree *Tree) insert(val int) {
	if tree == nil {
		tree = &Tree{
			Left:  nil,
			Value: val,
			Right: nil,
		}
		return
	}

	if tree.Left == nil {
			tree.Left = &Tree{
				Left: nil,
				Value: val,
				Right: nil,
			}
		return
	}
	if tree.Right == nil {
		tree.Right = &Tree{
			Left: nil,
			Value: val,
			Right: nil,
		}
		return
	}
	if tree.Left != nil {
		tree.Left.insert(val)
		return
	}
	if tree.Right != nil {
		tree.Right.insert(val)
		return
	}
	return
}

func (tree *Tree) print() {
	if tree == nil {
		fmt.Println("Empty")
		return
	}
	tree.Left.print()
	fmt.Println(tree.Value)
	tree.Right.print()
}

func main() {
	tree := &Tree{
		Value: 1,
	}
	tree.insert(10)
	tree.insert(2)
	tree.insert(6)
	tree.insert(9)
	tree.insert(4)
	tree.insert(15)
	tree.insert(20)
	tree.insert(3)
	tree.insert(7)
	tree.insert(12)
	tree.print()
}
