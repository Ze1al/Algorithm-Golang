package offer

import (
	"fmt"
	"testing"
)

func Test_findTargetSumWays(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3

	res := findTargetSumWays(nums, target)
	fmt.Println(res)
}
