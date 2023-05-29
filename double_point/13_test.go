package double_point

import (
	"fmt"
	"testing"
)

func Test_fourSumCount(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}

	res := fourSumCount(nums1, nums2, nums3, nums4)
	fmt.Printf("res=%v", res)
}
