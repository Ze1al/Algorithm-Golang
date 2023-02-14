package treenode

// inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 左中右                    左右中
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]
	root := &TreeNode{Val: val}

	if len(inorder) == 1 {
		return root
	}

	index := findIndex(inorder, val)
	root.Left = buildTree(inorder[:index], postorder[:len(inorder[:index])]) // 左闭右开  [)
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}

func findIndex(nums []int, x int) int {
	var res = -1
	for i, v := range nums {
		if v == x {
			res = i
			break
		}
	}
	return res
}
