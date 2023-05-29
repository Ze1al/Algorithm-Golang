package listnode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0

	dummy := &ListNode{Val: 0}
	p := dummy

	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}

		val := val1 + val2 + c

		p.Next = &ListNode{Val: val % 10}
		c = val / 10
	}

	if c > 0 {
		p.Next = &ListNode{Val: c}
	}
	return dummy.Next
}
