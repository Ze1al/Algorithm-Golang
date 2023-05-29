/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/24
 */

package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
