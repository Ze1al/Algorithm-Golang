/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/27
	反转链表
	定义两个前后指正，然后用倒罐子，最后同时横向移动
*/

package listnode

func reverseList(head *ListNode) *ListNode {
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
