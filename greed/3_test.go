package greed

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	res := maxSubArray(nums)
	fmt.Printf("res=%v", res)
}
