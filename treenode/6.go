package treenode

// 后序遍历迭代法
func postorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	cur := root
	stack := []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			res = append(res, cur.Val)
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			cur = cur.Right
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	reverse(res)
	return res
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
