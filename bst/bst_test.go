package bst

import "testing"

// Given a BST that looks like this:
//               8
//             /  \
//            4    9
//          /  \    \
//         2    5   10
//       /  \    \
//      1    3    6
//                 \
//                  7
// Test basic BST operations on it.

func TestBST(t *testing.T) {
	tree := &BST{}

	t.Run("Insert values into the tree", func(t *testing.T) {
		tree.Insert(8)
		tree.Insert(4)
		tree.Insert(9)
		tree.Insert(2)
		tree.Insert(5)
		tree.Insert(10)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(6)
		tree.Insert(7)
	})

	t.Run("Check the inorder traversal of the tree", func(t *testing.T) {
		nodes := tree.GetInorderTraversal()
		prev := nodes[0]
		for i := 1; i < len(nodes); i++ {
			cur := nodes[i]
			if cur <= prev {
				t.Errorf("Expected the inorder traversal of the tree to be ascending, found %d > %d", cur, prev)
				return
			}
			prev = nodes[i]
		}
	})

	t.Run("Check the leaf nodes of the tree", func(t *testing.T) {
		expectedLeafNodes := []int{1, 3, 7, 10}
		for _, nodeVal := range expectedLeafNodes {
			if node := tree.Search(nodeVal); node != nil && !node.IsLeaf() {
				t.Errorf("Expected %d to be a leaf node", nodeVal)
				return
			}
		}
	})

	t.Run("Delete leaf nodes from the tree", func(t *testing.T) {
		nodesToDelete := []int{3, 7, 10}
		newLeafNodes := []int{6, 9}

		for _, nodeVal := range nodesToDelete {
			tree.Delete(nodeVal)
			if deletedNode := tree.Search(nodeVal); deletedNode != nil {
				t.Errorf("Expected %d to be deleted from the tree", nodeVal)
				return
			}
		}

		for _, nodeVal := range newLeafNodes {
			if node := tree.Search(nodeVal); node != nil && !node.IsLeaf() {
				t.Errorf("Expected %d to be a new leaf node", nodeVal)
				return
			}
		}
	})

	t.Run("Delete nodes that has 1 child from the tree", func(t *testing.T) {
		nodesToDelete := []int{2, 5}
		for _, nodeVal := range nodesToDelete {
			tree.Delete(nodeVal)
			if deletedNode := tree.Search(nodeVal); deletedNode != nil {
				t.Errorf("Expected %d to be deleted from the tree", nodeVal)
				return
			}
		}

		node := tree.Search(4)
		if node.Left.Value != 1 {
			t.Error("Expected node 4's left child is now 1")
			return
		}
		if node.Right.Value != 6 {
			t.Error("Expected node 4's right child is now 6")
			return
		}
	})

	t.Run("Delete nodes that has 2 children from the tree", func(t *testing.T) {
		nodesToDelete := []int{4}
		for _, nodeVal := range nodesToDelete {
			tree.Delete(nodeVal)
			if deletedNode := tree.Search(nodeVal); deletedNode != nil {
				t.Errorf("Expected %d to be deleted from the tree", nodeVal)
				return
			}
		}

		node := tree.Search(1)
		if node.IsLeaf() {
			t.Error("Expected node 1 not to be a leaf node now")
			return
		}
		if node.Right.Value != 6 {
			t.Error("Expected node 1's right child is now 6")
			return
		}
	})

	t.Run("Delete the root node from the tree", func(t *testing.T) {
		tree.Delete(8)
		if node := tree.Search(6); node.Value != 6 {
			t.Error("Expected the new root value to be 6")
		}
	})

	t.Run("Delete all nodes from the tree", func(t *testing.T) {
		nodes := tree.GetInorderTraversal()
		for _, val := range nodes {
			tree.Delete(val)
		}

		if tree.root != nil {
			t.Errorf("Expected all nodes to be deleted from the tree, found root to be %d", tree.root.Value)
		}
	})
}
