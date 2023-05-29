package backTrace

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	nums := []int{1, 1, 2}
	res := permute(nums)
	fmt.Printf("res=%v", res)
}
