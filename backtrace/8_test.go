package backTrace

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}

	res := subsets(nums)
	fmt.Printf("res=%v", res)
}
