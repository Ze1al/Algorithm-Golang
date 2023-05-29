package listnode

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}

	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next

		temp := head.Next.Next
		head.Next.Next = head
		head.Next = temp

		pre = head

		head = temp
	}

	return dummy.Next
}
