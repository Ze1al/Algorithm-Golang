/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/25
	环的入口节点，如果没有没有就返回null

	解：f 快指针 s 慢指针
		f = 2s
		f相对于s来说就是多走了n圈 f = nb+s
		相减得到 f = 2nb s = nb
		入口节点： k = a + nb
		入口就是s节点多走了a
		第一次相遇以后，将fast节点放在入口，slow节点 s=s.Next f=f.Next 相遇就是入口节点
*/

package listnode

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 代表有环
			fast = head // 放回到head节点
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
