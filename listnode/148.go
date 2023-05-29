/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
	链表排序
*/

package listnode

import "sort"

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := []int{}
	cur := head
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}

	sort.Ints(list)
	newHead := &ListNode{Val: 0}
	p := newHead
	for i := 0; i < len(list); i++ {
		p.Next = &ListNode{Val: list[i]}
		p = p.Next
	}
	return newHead.Next
}
