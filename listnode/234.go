/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/28
	回文链表
*/

package listnode

// 加入到一个数组中，判断数组是不是回文数组
func IsPalindrome(head *ListNode) bool {
	list := []int{}
	cur := head
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}

	res := true
	for i, j := 0, len(list)-1; i < j; {
		if list[i] != list[j] {
			res = false
		}
		i++
		j--
	}
	return res
}

// 旋转链表来判断，旋转后面一半的链表
func _isPalindrome(head *ListNode) bool {
	mid := midListNode(head)
	half := reverselist(mid)

	p1 := head
	p2 := half
	res := true
	for p2 != nil && res {
		if p1.Val != p2.Val {
			res = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return res
}

func midListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverselist(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
