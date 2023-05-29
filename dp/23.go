package dp

func rob3(root *TreeNode) int {
	result := robTree(root)
	return max(result[0], result[1])
}

// 返回长度为2的数组，dp[0]代表不偷，dp[1]代表偷
func robTree(root *TreeNode) []int {
	var res = make([]int, 2)
	if root == nil {
		return res
	}

	left := robTree(root.Left)
	right := robTree(root.Right)

	// 偷root节点
	val1 := root.Val + left[0] + right[0]
	// 不偷root节点
	val2 := max(left[0], left[1]) + max(right[0], right[1])

	res = []int{val2, val1}
	return res
}
