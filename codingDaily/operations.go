package main

import (
	"fmt"
	"strconv"
)

type Tree struct {
	Value string
	Right *Tree
	Left  *Tree
}

func main() {

	tree := Tree{
		Value: "-",
		Left: &Tree{
			Value: "+",
			Left: &Tree{
				Value: "3",
				Left:  nil,
				Right: nil},
			Right: &Tree{
				Value: "-5",
				Left:  nil,
				Right: nil,
			}},
		Right: &Tree{
			Value: "+",
			Left: &Tree{
				Value: "4",
				Left:  nil,
				Right: nil,
			},
			Right: &Tree{
				Value: "3",
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(calculate(tree))
}

func calculate(t Tree) int {
	// In Order (left -> root -> right)
	if t.Left == nil && t.Right == nil {
		i, _ := strconv.Atoi(t.Value)
		return i
	}

	switch t.Value {
	case "+":
		return calculate(*t.Left) + calculate(*t.Right)
	case "-":
		return calculate(*t.Left) - calculate(*t.Right)
	case "*":
		return calculate(*t.Left) * calculate(*t.Right)
	default:
		return calculate(*t.Left) / calculate(*t.Right)
	}

}
