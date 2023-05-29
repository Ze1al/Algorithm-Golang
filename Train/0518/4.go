package _518

import (
	"strconv"
	"strings"
)

var res []string

// 二叉树的所有路劲
func binaryTreePaths(root *TreeNode) []string {
	res = []string{}

	backTravel(root, []string{})
	return res
}

func backTravel(root *TreeNode, path []string) {
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		res = append(res, strings.Join(path, "->"))
		return
	}

	if root.Left != nil {
		backTravel(root.Left, path)
	}
	if root.Right != nil {
		backTravel(root.Right, path)
	}
}
