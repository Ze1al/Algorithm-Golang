/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
	奇偶链表
*/

package listnode

func oddEvenList(head *ListNode) *ListNode {
	even := []int{}
	odd := []int{}

	cur := head
	flag := 1
	for cur != nil {
		if flag == 1 {
			even = append(even, cur.Val)
			flag = 0
		} else {
			odd = append(odd, cur.Val)
			flag = 1
		}
		cur = cur.Next
	}

	newList := &ListNode{Val: 0}
	p := newList
	for i := 0; i < len(even); i++ {
		p.Next = &ListNode{Val: even[i]}
		p = p.Next
	}
	for i := 0; i < len(odd); i++ {
		p.Next = &ListNode{Val: odd[i]}
		p = p.Next
	}
	return newList.Next
}
