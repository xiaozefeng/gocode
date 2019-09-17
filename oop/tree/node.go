package tree

type Node struct {
	Value       int
	Left, Right *Node
}

func (n *Node) Traversing(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.Traversing(f)
	f(n)
	n.Right.Traversing(f)
}

func (n *Node) TraversingWithChannel() chan *Node {
	var result = make(chan *Node)
	go func() {
		n.Traversing(func(node *Node) {
			result <- node
		})
		close(result)
	}()
	return result
}
