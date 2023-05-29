/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
 */

package listnode

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	occurred := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if !occurred[cur.Val] {
			occurred[cur.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}
