package double_point

import (
	"fmt"
	"testing"
)

func Test_threeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Printf("res=%v", res)
}
