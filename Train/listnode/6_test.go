package listnode

import (
	"fmt"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	headA := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	headB := &ListNode{Val: 1, Next: &ListNode{Val: 5}}

	node := getIntersectionNode(headA, headB)
	if node != nil {
		fmt.Printf("success")
	}
}
