package tree

type Node struct {
	value       int
	left, right *Node
}

func (n *Node) traversing(f func(*Node)) {
	if n == nil{
		return
	}
	n.left.traversing(f)
	f(n)
	n.right.traversing(f)
}
