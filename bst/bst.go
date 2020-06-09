package bst

// BST is a struct that provides operations
// on a Binary Search Tree.
type BST struct {
	root *TreeNode
}

// Insert adds a value to the BST.
func (t *BST) Insert(value int) {
	t.root = insert(t.root, value)
}

// Delete removes a value from the BST.
func (t *BST) Delete(value int) {
	t.root = delete(t.root, value)
}

// Search finds and returns the node which
// holds the specified value.
func (t *BST) Search(value int) *TreeNode {
	return search(t.root, value)
}

// GetInorderTraversal returns the inorder
// traversal of the BST.
func (t *BST) GetInorderTraversal() []int {
	return getNodesInorder(t.root)
}

func insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Value: value,
		}
	}

	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else if value > root.Value {
		root.Right = insert(root.Right, value)
	}

	return root
}

func delete(root *TreeNode, value int) *TreeNode {
	if root == nil || root.IsLeaf() {
		return nil
	}

	if value < root.Value {
		root.Left = delete(root.Left, value)
	} else if value > root.Value {
		root.Right = delete(root.Right, value)
	} else {
		if root.Left != nil {
			maxLeft := findMax(root.Left)
			root.Value = maxLeft
			root.Left = delete(root.Left, maxLeft)
		} else if root.Right != nil {
			minRight := findMin(root.Right)
			root.Value = minRight
			root.Right = delete(root.Right, minRight)
		}
	}

	return root
}

func search(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return nil
	}

	if value < root.Value {
		return search(root.Left, value)
	}

	if value > root.Value {
		return search(root.Right, value)
	}

	return root
}

func getNodesInorder(n *TreeNode) []int {
	if n.IsLeaf() {
		return []int{n.Value}
	}

	nodes := make([]int, 0)

	if n.Left != nil {
		leftNodes := getNodesInorder(n.Left)
		nodes = append(nodes, leftNodes...)
	}

	nodes = append(nodes, n.Value)

	if n.Right != nil {
		rightNodes := getNodesInorder(n.Right)
		nodes = append(nodes, rightNodes...)
	}

	return nodes
}
