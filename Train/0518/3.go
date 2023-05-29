package _518

// 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)

	cnt := left + right + 1

	return cnt
}
