package goroutine

import (
	"fmt"
	"github.com/xiaozefeng/gocode/oop/tree"
	"testing"
	"time"
)

func TestChannelDemo(t *testing.T) {
	channelDemo()
	time.Sleep(time.Microsecond)
}

func TestTraversingNodeWithChannel(t *testing.T) {
	var root = tree.Node{Value: 3}
	root.Left = &tree.Node{Value: 0}
	root.Left.Right = &tree.Node{Value: 2}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = &tree.Node{Value: 4}
	c := root.TraversingWithChannel()
	var maxValue int
	for n := range c {
		if n.Value > maxValue {
			maxValue = n.Value
		}
	}
	fmt.Println("max node value:", maxValue)

}
