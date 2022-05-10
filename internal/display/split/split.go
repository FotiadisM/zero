package split

import (
	"github.com/FotiadisM/zero/internal/display/view"
)

type SpliType int

const (
	SplitVert SpliType = iota
	SplitHorz
	SplitUndef
)

var leafIDCounter int

type Node struct {
	leafID int

	Type SpliType

	view.View

	Parent *Node

	Children [2]*Node
}

func newNode(id int, t SpliType, x, y, width, height int, parent *Node) *Node {
	v := view.View{X: x, Y: y, Width: width, Height: height}
	return &Node{
		leafID: id,
		Type:   t,
		View:   v,
		Parent: parent,
	}
}

func NewRootNode(x, y, width, height int) *Node {
	return newNode(0, SplitUndef, x, y, width, height, nil)
}

func (n *Node) GetLeafNode(id int) *Node {
	return nil
}

func (n *Node) HSplit(left bool) int {
	return 0
}

func (n *Node) VSplit(top bool) int {
	return 0
}

func (n *Node) UnSplit() {
}

func (n *Node) Resize() {}
