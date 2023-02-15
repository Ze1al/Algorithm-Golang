package dp

import (
	"fmt"
	"testing"
)

func Test_canPartition(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	res := canPartition(nums)
	fmt.Printf("res=%v", res)
}

func Test_canPartition2(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	res := canPartition(nums)
	fmt.Printf("res=%v", res)
}
