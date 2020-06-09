package bst

// TreeNode is a struct represents
// a node in a Binary Search Tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
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
