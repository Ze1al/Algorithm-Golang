package greed

import (
	"fmt"
	"testing"
)

func Test_canJump(t *testing.T) {
	nums := []int{1, 2, 3}
	res := canJump(nums)
	fmt.Printf("res=%v", res)
}
