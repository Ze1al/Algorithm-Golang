/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/4
	中 左 右
 */

package Travel

func preorder(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, preorder(root.Left)...)
	ans = append(ans, preorder(root.Right)...)
	return ans
}


func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	res := []int{}
	cur := root
	for cur != nil || len(stack) > 0{
		if cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur.Right)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
