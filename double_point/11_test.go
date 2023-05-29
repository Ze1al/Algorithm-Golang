package double_point

import (
	"fmt"
	"testing"
)

func Test_fourSum(t *testing.T) {
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	res := fourSum(nums, -11)
	fmt.Printf("res=%v", res)
}
