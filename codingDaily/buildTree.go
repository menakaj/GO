package main

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func main() {

	preOrder := []string{"a", "b", "d", "e", "c", "f", "g"}
	inOrder := []string{"d", "b", "e", "a", "f", "c", "g"}

	// 	    a
	//     /  \
	//    b     c
	//   / \    / \
	//  d   e  f   g
	construct(preOrder, inOrder)
}

func construct(pre []string, in []string) Node {
	tree := Node{
		Value: pre[0],
		Left:  nil,
		Right: nil,
	}
	return tree
}

func buildNode(left []string, right []string, value string) Node {
	if len(left) == 1 && len(right) == 1 {
		return Node{
			Value: value,
			Left:  &Node{Value: left[0], Left: nil, Right: nil},
			Right: &Node{Value: right[0], Left: nil, Right: nil},
		}
	}
	node := Node{
		Value: value,
		Left:  buildNode(),
	}
}
