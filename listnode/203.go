/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
	移除节点

*/

package listnode

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	cur := &ListNode{Val: 0, Next: head}

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
