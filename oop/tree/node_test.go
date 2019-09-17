package tree

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	var root = Node{Value: 3}
	root.Left = &Node{Value: 0}
	root.Left.Right = &Node{Value: 2}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = &Node{Value: 4}
}

func TestTraversing(t *testing.T) {

	var root = Node{Value: 3}
	root.Left = &Node{Value: 0}
	root.Left.Right = &Node{Value: 2}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = &Node{Value: 4}
	root.Traversing(func(n *Node) {
		fmt.Printf("%d ", n.Value)
	})
}
