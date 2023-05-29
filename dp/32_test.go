package dp

import (
	"fmt"
	"testing"
)

func Test_findLength(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 8}
	nums2 := []int{5, 6, 1, 4, 7}

	res := findLength(nums1, nums2)
	fmt.Printf("res=%d", res)
}
