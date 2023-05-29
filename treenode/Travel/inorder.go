/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/4
	左 中 右
 */

package Travel

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, inorderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)
	return ans
}

func inorder(root *TreeNode) []int{
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	res := []int{}
	cur := root
	for cur != nil || len(stack)>0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

