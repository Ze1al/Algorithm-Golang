package treenode

// 路径总和
// https://leetcode.cn/problems/path-sum/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	paths = [][]int{}
	findPath(root, []int{})
	for _, v := range paths {
		if sum(v) == targetSum {
			return true
		}
	}
	return false
}

func traversal(root *TreeNode, target int) bool {
	if target == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return false
	}

	if root.Left != nil {
		target -= root.Left.Val
		if traversal(root.Left, target) {
			return true
		}
		target += root.Left.Val
	}

	if root.Right != nil {
		target -= root.Right.Val
		if traversal(root.Right, target) {
			return true
		}
		target += root.Right.Val
	}
	return false
}

//var paths [][]int
//
//func findPath(root *TreeNode, path []int) {
//	path = append(path, root.Val)
//
//	if root.Left == nil && root.Right == nil {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		paths = append(paths, temp)
//		return
//	}
//	if root.Left != nil {
//		findPath(root.Left, path)
//		//path = path[:len(path)-1]
//	}
//	if root.Right != nil {
//		findPath(root.Right, path) // 有问题 没理解
//		//path = path[:len(path)-1]
//	}
//}
//
//
func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
