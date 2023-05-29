package listnode

// 移除重复的元素
// https://leetcode.cn/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	p := dummy

	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return dummy.Next
}
