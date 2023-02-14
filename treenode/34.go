package treenode

import (
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode
	minNum := math.MaxInt64

	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		if pre != nil {
			minNum = min(minNum, root.Val-pre.Val)
		}
		pre = root
		traversal(root.Right)
	}
	traversal(root)
	return minNum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
