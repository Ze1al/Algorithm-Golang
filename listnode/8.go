/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
 */

package listnode

func reverse1(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
