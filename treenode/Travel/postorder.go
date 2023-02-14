/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/4
	后序遍历
	左 右 中 -- > 中 右 左
 */

package Travel

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	ans = append(ans, postorderTraversal(root.Left)...)
	ans = append(ans, postorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}


func postorder(root *TreeNode) []int{
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	res := []int{}
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur.Left)
			cur = cur.Right
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return reverse(res)
}


func reverse(res []int) []int{
	for i, j := 0, len(res)-1; i<j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
