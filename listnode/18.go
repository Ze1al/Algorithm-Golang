/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/28
 */

package listnode

func deleteListNode(head *ListNode, val int) *ListNode {
	cur := head
	if head.Val == val {
		return head.Next
	}
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
