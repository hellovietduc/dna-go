package avl_tree

// TreeNode is a struct represents
// a node in an AVL Tree.
type TreeNode struct {
	Value  int
	Left   *TreeNode
	Right  *TreeNode
	height int
}

// IsLeaf returns whether this node
// is a leaf node.
func (n *TreeNode) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func findMin(n *TreeNode) int {
	if n.Left == nil {
		return n.Value
	}
	return findMin(n.Left)
}

func findMax(n *TreeNode) int {
	if n.Right == nil {
		return n.Value
	}
	return findMax(n.Right)
}

func (n *TreeNode) getBalanceFactor() int {
	if n.IsLeaf() {
		return 0
	}

	leftTreeHeight := -1
	rightTreeHeight := -1
	if n.Left != nil {
		leftTreeHeight = n.Left.height
	}
	if n.Right != nil {
		rightTreeHeight = n.Right.height
	}
	return leftTreeHeight - rightTreeHeight
}

func (n *TreeNode) rotateLeft() {
	rightNode := n.Right
	rightNodeLeftChild := rightNode.Left
	rightNodeRightChild := rightNode.Right

	newNode := &TreeNode{
		Value: n.Value,
		Left:  n.Left,
		Right: rightNodeLeftChild,
	}

	n.Value = rightNode.Value
	n.Left = newNode
	n.Right = rightNodeRightChild
}

func (n *TreeNode) rotateRight() {
	leftNode := n.Left
	leftNodeLeftChild := leftNode.Left
	leftNodeRightChild := leftNode.Right

	newNode := &TreeNode{
		Value: n.Value,
		Left:  leftNodeRightChild,
		Right: n.Right,
	}

	n.Value = leftNode.Value
	n.Left = leftNodeLeftChild
	n.Right = newNode
}
