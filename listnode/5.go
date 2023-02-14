/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/24
 */

package listnode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{}
	newHead := newList
	c := 0
	for l1 != nil || l2 != nil || c != 0 {
		var p1, p2 int
		if l1 != nil {
			p1 = l1.Val
		}
		if l2 != nil {
			p2 = l2.Val
		}

		sum := (p1 + p2 + c) % 10
		newHead.Next = &ListNode{Val: sum}
		c = (p1 + p2 + c) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		newHead = newHead.Next
	}
	return newList.Next
}
