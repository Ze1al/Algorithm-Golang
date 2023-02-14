/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/28
 */

package listnode

func numComponents(head *ListNode, G []int) int {
	dict := make(map[int]int)
	for _, v := range G {
		dict[v]++
	}

	res := 0
	flag := 0
	cur := head
	for cur != nil {
		if dict[cur.Val] > 0 {
			if flag == 0 {
				res++
				flag = 1
			}
		} else {
			flag = 0
		}
		cur = cur.Next
	}
	return res
}
