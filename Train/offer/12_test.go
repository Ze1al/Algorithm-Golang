package offer

import (
	"fmt"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	res := topKFrequent(nums, k)
	fmt.Println(res)
}
