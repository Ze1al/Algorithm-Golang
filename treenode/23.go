package treenode

// 高度平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if abs(right-left) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
