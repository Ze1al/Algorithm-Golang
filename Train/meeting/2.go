package meeting

// 百度
// 求N叉树的深度

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}

	var temp []int
	for _, node := range root.Children {
		temp = append(temp, maxDepth(node))
	}

	return maxArr(temp) + 1
}

func maxArr(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
