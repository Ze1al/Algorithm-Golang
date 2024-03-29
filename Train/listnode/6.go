package listnode

// 链表相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB

	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headA
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headB
		}
	}
	return p1
}
