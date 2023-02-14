package treenode

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/
// 二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var (
			sum    int
			length = len(queue)
		)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

			sum += cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		res = append(res, float64(sum)/float64(length))
	}
	return res
}
