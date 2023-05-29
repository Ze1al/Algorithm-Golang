package treenode

// 找到二叉树最底层 最左边的值

var (
	maxDepth2 int
	res       int
)

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	res, maxDepth2 = 0, 0

	backTracking2(root, 0)
	return res
}

func backTracking2(root *TreeNode, depth int) {
	if root.Left == nil && root.Right == nil {
		if depth > maxDepth2 {
			maxDepth2 = depth
			res = root.Val
		}
		return
	}

	if root.Left != nil {
		depth += 1
		backTracking2(root.Left, depth)
		depth -= 1
	}

	if root.Right != nil {
		depth += 1
		backTracking2(root.Right, depth)
		depth -= 1
	}
}
