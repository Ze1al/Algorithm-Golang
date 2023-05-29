package treenode

// 找到二叉树的最近公共祖先
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	return travel(root, p, q)
}

func travel(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
		left := travel(root.Left, p, q)
		if left != nil {
			return left
		}
	}

	if root.Val < p.Val && root.Val < q.Val {
		right := travel(root.Right, p, q)
		if right != nil {
			return right
		}
	}
	return root
}
