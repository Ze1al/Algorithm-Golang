package listnode

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	p := head
	for p != nil {
		temp := p.Next
		p.Next = pre
		pre = p
		p = temp
	}
	return pre
}
