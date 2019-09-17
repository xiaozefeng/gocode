package tree

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	var root = Node{value: 3}
	root.left = &Node{value: 0}
	root.left.right = &Node{value: 2}
	root.right = &Node{5, nil, nil}
	root.right.left = &Node{value: 4}
}

func TestTraversing(t *testing.T) {

	var root = Node{value: 3}
	root.left = &Node{value: 0}
	root.left.right = &Node{value: 2}
	root.right = &Node{5, nil, nil}
	root.right.left = &Node{value: 4}
	root.traversing(func(n *Node) {
		fmt.Printf("%d ", n.value)
	})
}
