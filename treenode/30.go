package treenode

// 最大二叉树
// https://leetcode.cn/problems/maximum-binary-tree/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	max, index := findMaxAndIndex(nums)
	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func findMaxAndIndex(nums []int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}

	index := 0
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}

	return max, index
}
