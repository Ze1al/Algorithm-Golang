package treenode

// 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var temp []int
		length := len(queue)

		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

			temp = append(temp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		if len(temp) > 0 {
			res = append(res, temp[len(temp)-1])
		}
	}
	return res
}
