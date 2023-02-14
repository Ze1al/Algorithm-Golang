package treenode

// 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth int

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		depth += 1

		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return depth
}
