package backTrace

import (
	"fmt"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	nums := []int{4, 6, 7, 7}

	res := findSubsequences(nums)
	fmt.Printf("res=%v", res)
}
