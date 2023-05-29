/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/30
	二进制链表转整数
*/

package listnode

func getDecimalValue(head *ListNode) int {
	var res int
	if head == nil {
		return res
	}
	cur := head

	for cur != nil {
		res = res*2 + cur.Val
		cur = cur.Next
	}
	return res
}
