package treenode

// 递归后序遍历 https://leetcode.cn/problems/binary-tree-postorder-traversal/
// 左右中 反过来就是中右左
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	left := postorderTraversal(root.Left)
	res = append(res, left...)
	right := postorderTraversal(root.Right)
	res = append(res, right...)

	res = append(res, root.Val)
	return res
}
