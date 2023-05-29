/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/24
	返回倒数第k个节点
*/

package listnode

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return 0
	}
	fast := head
	slow := head

	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}
