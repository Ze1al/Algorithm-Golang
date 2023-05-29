package offer

import (
	"fmt"
	"testing"
)

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}

	res := containsNearbyAlmostDuplicate(nums, 3, 0)
	fmt.Println(res)
}
