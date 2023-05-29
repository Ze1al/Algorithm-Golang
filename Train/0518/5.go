package _518

// 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil && targetSum == 0 {
		return true
	}

	return backTravel2(root, targetSum)
}

func backTravel2(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && target == root.Val {
		return true
	}

	if root.Left != nil {
		target -= root.Val
		if backTravel2(root.Left, target) {
			return true
		}
		target += root.Val
	}

	if root.Right != nil {
		target -= root.Val
		if backTravel2(root.Right, target) {
			return true
		}
		target += root.Val
	}
	return false
}
