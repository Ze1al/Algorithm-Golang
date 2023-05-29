package offer

import (
	"fmt"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	nums := []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}
	target := 15

	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
