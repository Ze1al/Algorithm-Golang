package treenode

// 找出每一行中最大值
// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		temp := []int{}
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

		res = append(res, findMax(temp))
	}
	return res
}

func findMax(num []int) int {
	if len(num) == 0 {
		return 0
	}

	max := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
		}
	}
	return max
}
