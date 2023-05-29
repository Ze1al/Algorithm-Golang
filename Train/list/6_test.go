package list

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2, 7}

	res := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("res=%f", res)
}
