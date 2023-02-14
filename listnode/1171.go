/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
	删除总和为零的连续节点
*/

package listnode

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	list := []int{}
	cur := head
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}
	listNode := deleteSumEqualZero(list)
	newHead := &ListNode{Val: 0}
	p := newHead
	for i := 0; i < len(listNode); i++ {
		p.Next = &ListNode{Val: list[i]}
		p = p.Next
	}
	return newHead.Next
}

func deleteSumEqualZero(list []int) []int {
	for i := 0; i < len(list); i++ {
		sum := list[i]
		for j := i + 1; j < len(list); j++ {
			sum += list[j]
			if sum == 0 {
				list = list[i:]
				list = append(list, list[:j]...)
			}
		}
	}
	return list
}
