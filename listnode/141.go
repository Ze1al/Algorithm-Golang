/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/25
	判断链表是否有环
*/

package listnode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
