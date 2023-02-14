package treenode

// 找到最近的公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归终止条件
	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil && right == nil {
		return left
	}

	if right != nil && left == nil {
		return right
	}
	return nil
}
