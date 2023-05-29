package _518

// 中序遍历
// 左 中 右
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := []*TreeNode{}
	cur := root
	for len(stack) > 0 || cur != nil {
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
