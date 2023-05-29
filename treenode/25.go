package treenode

// 左只树之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	left := sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		left += root.Left.Val
	}

	right := sumOfLeftLeaves(root.Right)

	sum := left + right
	return sum
}
