package meeting

// 百度
// 求二叉树的深度

// 求深度，其实是从下往上，最后一个节点的深度为1
// 所以是 左右中，后序遍历
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)
	return max(left, right) + 1
}
