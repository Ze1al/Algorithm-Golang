/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
	合并两个链表
[0,1,2,3,4,5]
3
4
[1000000,1000001,1000002]
*/

package listnode

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	p1 := list1
	p2 := list2

	for p1 != nil {
		if p1.Val == a {
			p1.Next = p2
			for p2.Next != nil {
				p2 = p2.Next
			}
			cur := p1
			for cur != nil && cur.Val == b {
				cur = cur.Next
			}
			cur.Next = p2

		} else {
			p1 = p1.Next
		}
	}
	return list1
}
