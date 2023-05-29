package treenode

// 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (right == nil && left != nil) || (left.Val != right.Val) {
		return false
	}
	if !(symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)) {
		return false
	}

	return true
}
