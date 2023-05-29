package dp

import (
	"fmt"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(nums)
	fmt.Printf("res=%d", res)
}
