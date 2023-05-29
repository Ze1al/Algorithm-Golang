package interview

import (
	"fmt"
	"testing"
)

func Test_threeSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 14

	res := threeSum(nums, n)
	fmt.Println(res)
}
