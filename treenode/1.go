package treenode

// 递归前序遍历  中左右
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	left := preorderTraversal(root.Left)
	res = append(res, left...)
	right := preorderTraversal(root.Right)
	res = append(res, right...)
	return res
}
