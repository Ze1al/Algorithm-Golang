package treenode

// 从前序遍历和中序遍历 构造二叉树
// 中 左 右			左 中 右
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	root := &TreeNode{Val: val}
	if len(preorder) == 1 {
		return root
	}
	index := findIndex(inorder, val)

	root.Left = buildTree2(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree2(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}
