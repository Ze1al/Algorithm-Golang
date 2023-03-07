package monotonic_stack

import (
	"fmt"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	res := nextGreaterElement(nums1, nums2)
	fmt.Printf("res=%v", res)
}
