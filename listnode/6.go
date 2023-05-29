/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/24
	回文字符串
*/

package listnode

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	mid := findMid(head)
	halfListNode := reverseList(mid.Next)
	p1 := head
	p2 := halfListNode
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	mid.Next = reverseList(halfListNode)
	return true
}

func findMid(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

/*
	反转链表
	定义两个指针
	pre 和 cur实现一次局部的反转，然后pre += 1， last += 1
*/
func reverse(head *ListNode) *ListNode {
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
