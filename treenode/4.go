package treenode

// 前序遍历
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return res
}

// 前序遍历 中左右
func preTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	cur := root
	stack := []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			res = append(res, cur.Val)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
