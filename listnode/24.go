/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/27
	两两交换链表中的节点
*/

package listnode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		pre := cur.Next
		last := cur.Next.Next
		cur.Next = last
		pre.Next = last.Next
		last.Next = pre
		cur = pre
	}
	return dummy.Next
}
