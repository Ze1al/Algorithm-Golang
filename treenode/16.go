package treenode

// 链接下一个节点
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		length := len(queue)
		temp := []*Node{}

		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

			temp = append(temp, cur)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		for i := 0; i < len(temp)-1; i++ {
			temp[i].Next = temp[i+1]
		}
	}
	return root
}
