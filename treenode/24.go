package treenode

import "strconv"

// 二叉树的所有路径
// https://leetcode.cn/problems/binary-tree-paths/

var paths [][]int

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	paths = [][]int{}
	findPath(root, []int{})
	for _, v := range paths {
		var s string
		for i, val := range v {
			s += strconv.Itoa(val)
			if i != len(v)-1 {
				s += "->"
			}
		}

		res = append(res, s)
	}
	return res
}

func findPath(root *TreeNode, path []int) {
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil {
		temp := make([]int, len(path))
		copy(temp, path)
		paths = append(paths, temp)
		return
	}
	if root.Left != nil {
		findPath(root.Left, path)
		//path = path[:len(path)-1]
	}
	if root.Right != nil {
		findPath(root.Right, path) // 有问题 没理解
		//path = path[:len(path)-1]
	}
}
