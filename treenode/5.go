package treenode

// 迭代版 中序遍历
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	cur := root
	stack := []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
