/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
	获取中间节点
*/

package listnode

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
