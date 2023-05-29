package listnode

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_mergeListNode(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}

	res := mergeListNode(l1, l2)
	s, _ := json.Marshal(res)
	fmt.Printf("res=%s", string(s))
}
