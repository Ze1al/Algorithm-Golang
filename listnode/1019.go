/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
 */

package listnode

func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{0}
	}

	res := []int{}
	slow := head
	for slow != nil && slow.Next != nil {
		fast := slow.Next
		temp := 0
		for fast != nil {
			if fast.Val > slow.Val && fast.Val > temp {
				temp = fast.Val
				break
			}
			fast = fast.Next
		}
		slow = slow.Next
		res = append(res, temp)
	}
	res = append(res, 0)
	return res
}
