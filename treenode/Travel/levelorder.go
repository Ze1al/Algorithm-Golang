/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/4
	层序遍历
 */

package Travel

func levelorder(root *TreeNode) []int {
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) >0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		res = append(res, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return res
}


func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		ret = append(ret, []int{})
		temp := []*TreeNode{}
		for j:=0; j<len(queue); j++ {
			cur := queue[j]
			ret[i] = append(ret[i], cur.Val)
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
		}
		queue = temp
	}
	return ret
}






