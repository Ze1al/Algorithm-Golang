package treenode

// 中序遍历 左中右 递归遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	left := inorderTraversal(root.Left)
	res = append(res, left...)
	res = append(res, root.Val)
	right := inorderTraversal(root.Right)
	res = append(res, right...)

	return res
}
