/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/24
	分割链表
	使用两个链表
	左边小于x
	右边大于x
*/

package listnode

func partition(head *ListNode, x int) *ListNode {
	cur := head
	left := &ListNode{
		Val: 0,
	}
	right := &ListNode{
		Val: 0,
	}
	leftP := left
	rightP := right
	for cur != nil {
		if cur.Val < x {
			left.Next = &ListNode{Val: cur.Val}
			left = left.Next
		} else {
			right.Next = &ListNode{Val: cur.Val}
			right = right.Next
		}
		cur = cur.Next
	}

	leftP = leftP.Next
	rightP = rightP.Next

	if leftP == nil {
		return rightP
	}

	res := leftP
	for leftP.Next != nil {
		leftP = leftP.Next
	}
	left.Next = rightP

	return res
}
