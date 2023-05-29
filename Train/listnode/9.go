package listnode

// 合并两个有序列表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	p := dummy

	l1 := list1
	l2 := list2

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			p.Next = &ListNode{Val: l2.Val}
			p = p.Next

			l2 = l2.Next
		} else {
			p.Next = &ListNode{Val: l1.Val}
			p = p.Next

			l1 = l1.Next
		}
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return dummy.Next
}
