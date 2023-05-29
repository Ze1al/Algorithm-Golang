package treenode

// 二叉树删除节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	minnode := root.Right
	for minnode.Left != nil {
		minnode = minnode.Left
	}
	root.Val = minnode.Val
	root.Right = deleteNode1(root.Right)
	return root
}

func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}
