package avl_tree

// AVLTree is a struct that provides operations
// on an AVL Tree.
type AVLTree struct {
	root *TreeNode
}

// Insert adds a value to the AVL Tree.
func (t *AVLTree) Insert(value int) {
	t.root = insert(t.root, value)
}

// Delete removes a value from the AVL Tree.
func (t *AVLTree) Delete(value int) {
	t.root = delete(t.root, value)
}

// Search finds and returns the node which
// holds the specified value.
func (t *AVLTree) Search(value int) *TreeNode {
	return search(t.root, value)
}

// GetInorderTraversal returns the inorder
// traversal of the AVL Tree.
func (t *AVLTree) GetInorderTraversal() []int {
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
		updateHeight(root, true)
		ensureBalance(root)
	} else if value > root.Value {
		root.Right = insert(root.Right, value)
		updateHeight(root, false)
		ensureBalance(root)
	}

	return root
}

func delete(root *TreeNode, value int) *TreeNode {
	if root == nil || root.IsLeaf() {
		return nil
	}

	if value < root.Value {
		root.Left = delete(root.Left, value)
		updateHeight(root, true)
		ensureBalance(root)
	} else if value > root.Value {
		root.Right = delete(root.Right, value)
		updateHeight(root, false)
		ensureBalance(root)
	} else {
		if root.Left != nil {
			maxLeft := findMax(root.Left)
			root.Value = maxLeft
			root.Left = delete(root.Left, maxLeft)
			updateHeight(root, true)
			ensureBalance(root)
		} else if root.Right != nil {
			minRight := findMin(root.Right)
			root.Value = minRight
			root.Right = delete(root.Right, minRight)
			updateHeight(root, false)
			ensureBalance(root)
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

func updateHeight(node *TreeNode, isLeft bool) {
	leftHeight := -1
	rightHeight := -1

	if node.Left != nil {
		leftHeight = node.Left.height
	}
	if node.Right != nil {
		rightHeight = node.Right.height
	}

	if isLeft {
		leftHeight++
	} else {
		rightHeight++
	}

	node.height = max(leftHeight, rightHeight)
}

func ensureBalance(node *TreeNode) {
	bf := node.getBalanceFactor()
	if bf >= -1 && bf <= 1 {
		return
	}

	if bf > 1 {
		// tree is left-heavy
		if node.Left.Left == nil {
			node.Left.rotateLeft() // left-right rotation
		}
		node.rotateRight()
	} else {
		// tree is right-heavy
		if node.Right.Right == nil {
			node.Right.rotateRight() // right-left rotation
		}
		node.rotateLeft()
	}
}
