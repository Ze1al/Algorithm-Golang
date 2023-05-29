/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/27
	删除链表中的重复元素
*/

package listnode

import "sort"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dict := make(map[int]int)
	cur := head
	for cur != nil {
		if _, ok := dict[cur.Val]; ok {
			dict[cur.Val]++
		} else {
			dict[cur.Val] = 1
		}
		cur = cur.Next
	}

	list := []int{}
	for k, _ := range dict {
		list = append(list, k)
	}
	sort.Ints(list)

	dummy := &ListNode{Val: 0}
	current := dummy
	for i := 0; i < len(list); i++ {
		if dict[list[i]] == 1 {
			current.Next = &ListNode{Val: list[i]}
			current = current.Next
		}
	}
	return dummy.Next
}
